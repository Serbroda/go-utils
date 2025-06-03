//go:build !windows

package process

import (
	"os/exec"
	"syscall"
)

func IsRunning(cmd *exec.Cmd) bool {
	if cmd == nil || cmd.Process == nil {
		return false
	}
	err := cmd.Process.Signal(syscall.Signal(0))
	return err == nil
}
