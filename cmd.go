package input_autocomplete

import (
	"io/ioutil"
	"os"

	"golang.org/x/sys/windows"
)

type Cmd struct{}

type DirLister interface {
	ListContent(path string) ([]string, error)
}

type DirChecker interface {
	IsDir(path string) (bool, error)
}

type DirListChecker interface {
	DirLister
	DirChecker
}

func (c Cmd) ListContent(path string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}

	return files, nil
}

func (c Cmd) IsDir(path string) (bool, error) {
	info, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}

func EnableVirtalTerminalWindows() error {
	var originalMode uint32
	stdout := windows.Handle(os.Stdout.Fd())

	if err := windows.GetConsoleMode(stdout, &originalMode); err != nil {
		return err
	}

	return windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
