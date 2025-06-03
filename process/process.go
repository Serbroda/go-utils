package process

import "os/exec"

func RunAsync(cmd *exec.Cmd, callback func(err error)) {
	go func() {
		err := cmd.Run()
		callback(err)
	}()
}
