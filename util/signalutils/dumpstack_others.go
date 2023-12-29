//go:build !windows
// +build !windows

package signalutils

import (
	"syscall"

	"github.com/nyl1001/pkg/utils"
	"yunion.io/x/log"
)

func SetDumpStackSignal() {
	RegisterSignal(func() {
		utils.DumpAllGoroutineStack(log.Logger().Out)
	}, syscall.SIGUSR1)
}
