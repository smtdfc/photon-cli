package helpers

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func SpawnCommand(name string, args []string, inheritIO bool) (int, error) {
	cmd := exec.Command(name, args...)

	if inheritIO {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	} else {
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		go io.Copy(os.Stdout, stdout)
		go io.Copy(os.Stderr, stderr)
	}

	err := cmd.Start()
	if err != nil {
		return -1, fmt.Errorf("Cannot spawn command: %w", err)
	}

	err = cmd.Wait()
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode(), err
	} else if err != nil {
		return -1, err
	}

	return 0, nil
}
