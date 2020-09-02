package input_autocomplete

import (
	"runtime"
	"strings"
)

type autocomplete struct {
	cmd DirLister
}

func Autocomplete(path string) (string, error) {
	os := runtime.GOOS
	switch os {
	case "linux":
		a := autocomplete{
			cmd: CmdLinux{},
		}
		return a.linuxAutocomplete(path)
	case "darwin":
		return path, nil
	case "windows":
		return path, nil
	default:
		return path, nil
	}
}

func (a autocomplete) linuxAutocomplete(path string) (string, error) {
	var splittedPath []string
	if path[0] == '/' {
		splittedPath = strings.Split(path[1:], "/")
	} else {
		splittedPath = strings.Split(path, "/")
	}

	lastValidSplittedPath := splittedPath[:len(splittedPath)-1]

	var lastValidPath string
	for _, subPath := range lastValidSplittedPath {
		lastValidPath += "/" + subPath
	}

	contents, err := a.cmd.ListContent(lastValidPath)
	if err != nil {
		return path, err
	}

	for _, dir := range contents {
		if strings.HasPrefix(dir, splittedPath[len(splittedPath)-1]) {
			newPath := append(lastValidSplittedPath, dir)
			return "/" + strings.Join(newPath, "/") + "/", nil
		}
	}

	return path, nil
}
