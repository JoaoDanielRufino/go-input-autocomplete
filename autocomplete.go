package input_autocomplete

import (
	"os"
	"runtime"
	"strings"
)

type autocomplete struct {
	cmd DirLister
}

func Autocomplete(path string) string {
	os := runtime.GOOS
	switch os {
	case "linux", "darwin":
		a := autocomplete{
			cmd: CmdUnix{},
		}
		return a.unixAutocomplete(path)
	case "windows":
		return path
	default:
		return path
	}
}

// Return if the string starts with prefix, case insensitive
func hasInsensitivePrefix(s string, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[0:len(prefix)], prefix)
}

func (a autocomplete) unixAutocomplete(path string) string {
	if path == "" || path[len(path)-1] == ' '{
		return path
	}
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash == -1 || (path[0] != '/' && path[:2] != "./"){
		path = "./" + path
		lastSlash = 1
	}
	path = a.findFromPrefix(path, lastSlash)
	ok, err := isDir(path)
	if ok && err == nil {
		path = path + "/"
	}
	return path
}

func (a autocomplete) findFromPrefix(prefix string, lastSlash int) string {
	contents, err := a.cmd.ListContent(prefix[:lastSlash+1])
	if err != nil {
		return prefix
	}
	for _, content := range contents {
		if hasInsensitivePrefix(content, prefix[lastSlash+1:]) {
			return prefix[:lastSlash+1] + content
		}
	}
	return prefix
}

func isDir(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
