package input_autocomplete

import (
	"runtime"
	"strings"
)

type autocomplete struct {
	cmd DirListChecker
}

func Autocomplete(path string) []string {
	os := runtime.GOOS
	a := autocomplete{
		cmd: Cmd{},
	}

	switch os {
	case "linux", "darwin":
		return a.unixAutocomplete(path)
	case "windows":
		return a.windowsAutocomplete(path)
	default:
		return []string{path}
	}
}

func invalidPath(path string) bool {
	if path == "" || strings.TrimSpace(path) == "" {
		return true
	}

	return false
}

// Return if the string starts with prefix, case insensitive
func hasInsensitivePrefix(s string, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[0:len(prefix)], prefix)
}

func (a autocomplete) unixAutocomplete(path string) []string {
	if invalidPath(path) {
		return []string{path}
	}

	lastSlash := strings.LastIndex(path, "/")
	if lastSlash == -1 || (path[0] != '/' && path[:2] != "./") {
		path = "./" + path
		if !strings.Contains(path, "/") {
			lastSlash = 1
		} else {
			lastSlash = strings.LastIndex(path, "/")
		}
	}

	return a.findFromPrefix(path, lastSlash, "/")
}

func (a autocomplete) windowsAutocomplete(path string) []string {
	if invalidPath(path) {
		return []string{path}
	}

	lastSlash := strings.LastIndex(path, "\\")
	if !strings.Contains(path, ":") && !strings.Contains(path, ".\\") {
		path = ".\\" + path
		if !strings.Contains(path, "\\") {
			lastSlash = 1
		} else {
			lastSlash = strings.LastIndex(path, "\\")
		}
	}

	return a.findFromPrefix(path, lastSlash, "\\")
}

func (a autocomplete) findFromPrefix(prefix string, lastSlash int, sep string) []string {
	contents, err := a.cmd.ListContent(prefix[:lastSlash+1])
	if err != nil {
		return []string{prefix}
	}

	var matches []string
	for _, content := range contents {
		if hasInsensitivePrefix(content, prefix[lastSlash+1:]) {
			p := prefix[:lastSlash+1] + content
			ok, err := a.cmd.IsDir(p)
			if ok && err == nil {
				p = p + sep
			}
			matches = append(matches, p)
		}
	}
	if len(matches) == 0 {
		matches = append(matches, prefix)
	}

	return matches
}
