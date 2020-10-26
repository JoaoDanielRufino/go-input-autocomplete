package input_autocomplete

import (
	"io/ioutil"
)

type CmdUnix struct{}

type CmdWindows struct{}

type DirLister interface {
	ListContent(path string) ([]string, error)
}

func (c CmdUnix) ListContent(path string) ([]string, error) {
	return readDir(path)
}

// readDir reads the directory named by root and
// returns a list of directory entries sorted by filename.
func readDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
