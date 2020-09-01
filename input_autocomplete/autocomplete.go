package input_autocomplete

import (
	"fmt"
	"runtime"
)

type autocomplete struct {
	cmd DirLister
}

func Autocomplete(text string) (string, error) {
	os := runtime.GOOS
	switch os {
	case "linux":
		a := autocomplete{
			cmd: CmdLinux{},
		}
		return a.linuxAutocomplete(text)
	case "darwin":
		return text, nil
	case "windows":
		return text, nil
	default:
		return text, nil
	}
}

func (a autocomplete) linuxAutocomplete(text string) (string, error) {
	contents, _ := a.cmd.ListContent(text)

	fmt.Printf("\n%q\n", contents)

	return text, nil
}
