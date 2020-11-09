package input_autocomplete

import (
	"io/ioutil"
)

type Cmd struct{}

type DirLister interface {
	ListContent(path string) ([]string, error)
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
