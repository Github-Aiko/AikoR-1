package cmd

import (
	"errors"
	"os"
	"os/exec"
)

func execCommand(cmd string) (string, error) {
	e := exec.Command("bash", "-c", cmd)
	out, err := e.CombinedOutput()
	if errors.Unwrap(err) == exec.ErrNotFound {
		e = exec.Command("sh", "-c", cmd)
		out, err = e.CombinedOutput()
	}
	return string(out), err
}

func execCommandStd(name string, args ...string) {
	e := exec.Command(name, args...)
	e.Stdout = os.Stdout
	e.Stdin = os.Stdin
	e.Stderr = os.Stderr
	_ = e.Run()
}