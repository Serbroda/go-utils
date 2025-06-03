//go:build process_test

package process

import (
	"os/exec"
	"runtime"
	"testing"
	"time"
)

func TestRunAsync_Success(t *testing.T) {
	done := make(chan error, 1)

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "ping", "127.0.0.1", "-n", "2") // 2x ping
	} else {
		cmd = exec.Command("sleep", "1") // 1 Sekunde schlafen
	}

	RunAsync(cmd, func(err error) {
		done <- err
	})

	select {
	case err := <-done:
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Error("timeout waiting for async command to finish")
	}
}

func TestIsRunning(t *testing.T) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "127.0.0.1", "-n", "5")
	} else {
		cmd = exec.Command("sleep", "2")
	}

	err := cmd.Start()
	if err != nil {
		t.Fatalf("failed to start process: %v", err)
	}
	defer cmd.Process.Kill() // sauber aufräumen, falls was schiefläuft

	if !IsRunning(cmd) {
		t.Error("expected process to be running")
	}

	// Warte, bis Prozess von selbst endet
	err = cmd.Wait()
	if err != nil {
		t.Errorf("process exited with error: %v", err)
	}

	if IsRunning(cmd) {
		t.Error("expected process to have stopped")
	}
}
