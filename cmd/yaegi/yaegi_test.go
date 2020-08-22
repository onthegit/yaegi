package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

const (
	// CITimeoutMultiplier is the multiplier for all timeouts in the CI.
	CITimeoutMultiplier = 3
)

// Sleep pauses the current goroutine for at least the duration d.
func Sleep(d time.Duration) {
	d = applyCIMultiplier(d)
	time.Sleep(d)
}

func applyCIMultiplier(timeout time.Duration) time.Duration {
	ci := os.Getenv("CI")
	if ci == "" {
		return timeout
	}
	b, err := strconv.ParseBool(ci)
	if err != nil || !b {
		return timeout
	}
	return time.Duration(float64(timeout) * CITimeoutMultiplier)
}

func TestYaegiCmdCancel(t *testing.T) {
	tmp, err := ioutil.TempDir("", "yaegi-")
	if err != nil {
		t.Fatalf("failed to create tmp directory: %v", err)
	}
	defer func() {
		err = os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("failed to clean up %v: %v", tmp, err)
		}
	}()

	yaegi := filepath.Join(tmp, "yaegi")
	build := exec.Command("go", "build", "-race", "-o", yaegi, ".")
	err = build.Run()
	if err != nil {
		t.Fatalf("failed to build yaegi command: %v", err)
	}

	// Test src must be terminated by a single newline.
	tests := []string{
		"for {}\n",
		"select {}\n",
	}
	for _, src := range tests {
		cmd := exec.Command(yaegi)
		in, err := cmd.StdinPipe()
		if err != nil {
			t.Errorf("failed to get stdin pipe to yaegi command: %v", err)
		}
		var outBuf, errBuf bytes.Buffer
		cmd.Stdout = &outBuf
		cmd.Stderr = &errBuf

		// https://golang.org/doc/articles/race_detector.html#Options
		cmd.Env = []string{`GORACE="halt_on_error=1"`}

		err = cmd.Start()
		if err != nil {
			t.Fatalf("failed to start yaegi command: %v", err)
		}

		_, err = in.Write([]byte(src))
		if err != nil {
			t.Errorf("failed pipe test source to yaegi command: %v", err)
		}
		Sleep(200 * time.Millisecond)
		err = cmd.Process.Signal(os.Interrupt)
		if err != nil {
			t.Errorf("failed to send os.Interrupt to yaegi command: %v", err)
		}

		_, err = in.Write([]byte("1+1\n"))
		if err != nil {
			t.Errorf("failed to probe race: %v", err)
		}
		err = in.Close()
		if err != nil {
			t.Errorf("failed to close stdin pipe: %v", err)
		}

		err = cmd.Wait()
		if err != nil {
			if cmd.ProcessState.ExitCode() == 66 { // See race_detector.html article.
				t.Errorf("race detected running yaegi command canceling %q: %v", src, err)
				if testing.Verbose() {
					t.Log(&errBuf)
				}
			} else {
				t.Errorf("error running yaegi command for %q: %v", src, err)
			}
			continue
		}

		if outBuf.String() != "context canceled\n" {
			t.Errorf("unexpected output: %q", &outBuf)
		}
	}
}
