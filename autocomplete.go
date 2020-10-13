package input_autocomplete

import (
	"os"
	"runtime"
	"strings"
)

type autocomplete struct {
	cmd DirLister
}

func Autocomplete(path string) (string, error) {
	os := runtime.GOOS
	switch os {
	case "linux", "darwin":
		a := autocomplete{
			cmd: CmdUnix{},
		}
		return a.unixAutocomplete(path)
	case "windows":
		return path, nil
	default:
		return path, nil
	}
}

// Return if the string starts with prefix, case insensitive
func hasInsensitivePrefix(s string, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[0:len(prefix)], prefix)
}

func (a autocomplete) unixAutocomplete(path string) (string, error) {
	if path == "" {
		return "", nil
	}
	var splittedPath []string
	if path[0] == '/' {
		splittedPath = strings.Split(path[1:], "/")
	} else {
		splittedPath = strings.Split(path, "/")
	}

	lastValidSplittedPath := splittedPath[:len(splittedPath)-1]

	lastValidPath := strings.Join(lastValidSplittedPath, "/")

	if lastValidPath == "" {
		if path == "." {
			lastValidPath = "."
		} else {
			lastValidPath = "/"
		}
	} else if lastValidPath[0] != '.' {
		lastValidPath = "/" + lastValidPath
	}

	if !isDir(lastValidPath) {
		return lastValidPath, nil
	} else if len(lastValidPath) != 0 {
		lastValidPath += "/"
	}

	contents, err := a.cmd.ListContent(lastValidPath)
	if err != nil {
		return lastValidPath, err
	}

	for _, dir := range contents {
		if hasInsensitivePrefix(dir, splittedPath[len(splittedPath)-1]) {
			newPathSlice := append(lastValidSplittedPath, dir)
			newPath := strings.Join(newPathSlice, "/")
			if newPath[0] != '.' {
				newPath = "/" + newPath
			}
			if isDir(newPath) {
				newPath += "/"
			}
			return newPath, nil
		}
	}
	return lastValidPath, nil
}

func isDir(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return info.IsDir()
}
