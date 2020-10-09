package input_autocomplete

import (
	"os/exec"
	"strings"
)

type CmdUnix struct{}

type CmdWindows struct{}

type DirLister interface {
	ListContent(path string) ([]string, error)
}

func (c CmdUnix) ListContent(path string) ([]string, error) {
	cmd := exec.Command("ls", "-a", path)
	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lsOutput := strings.Split(string(stdout), "\n")

	return lsOutput, nil
}
