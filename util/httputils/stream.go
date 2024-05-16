package httputils

import (
	"bytes"
	"io"
	"mime/multipart"
)

func ReadFileForm(r *multipart.Reader) (map[string]string, *multipart.Part, error) {
	params := make(map[string]string)
	maxValueBytes := int64(10 << 20)
	for {
		p, err := r.NextPart()
		if err == io.EOF {
			return params, nil, nil
		}
		if err != nil {
			return nil, nil, err
		}

		name := p.FormName()
		if name == "" {
			continue
		}
		filename := p.FileName()

		var b bytes.Buffer

		_, hasContentTypeHeader := p.Header["Content-Type"]
		if !hasContentTypeHeader && filename == "" {
			// value, store as string in memory
			n, err := io.CopyN(&b, p, maxValueBytes+1)
			if err != nil && err != io.EOF {
				return nil, nil, err
			}
			maxValueBytes -= n
			if maxValueBytes < 0 {
				return nil, nil, multipart.ErrMessageTooLarge
			}
			params[name] = b.String()
			continue
		}

		if name == "image" || name == "file" {
			params["_file_name"] = filename
			return params, p, nil
		} else {
			//return params, nil, fmt.Errorf("no file uploaded")
			return params, nil, nil
		}
	}
}
