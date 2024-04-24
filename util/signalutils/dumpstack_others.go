//go:build !windows
// +build !windows

package signalutils

import (
	"syscall"

	"github.com/nyl1001/log"
	"github.com/nyl1001/pkg/utils"
)

func SetDumpStackSignal() {
	RegisterSignal(func() {
		utils.DumpAllGoroutineStack(log.Logger().Out)
	}, syscall.SIGUSR1)
}
