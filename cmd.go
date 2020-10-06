package input_autocomplete

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CmdLinux struct{}

type CmdDarwin struct{}

type CmdWindows struct{}

type DirUtil interface {
	ListContent(path string) ([]string, error)
	IsDir(path string) bool
}

func (c CmdLinux) ListContent(path string) ([]string, error) {
	cmd := exec.Command("ls", path)
	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lsOutput := strings.Split(string(stdout), "\n")

	return lsOutput, nil
}

func (c CmdLinux) IsDir(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return info.IsDir()
}
