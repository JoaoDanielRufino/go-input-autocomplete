package input_autocomplete

import "runtime"

func Autocomplete(text string) (string, error) {
	os := runtime.GOOS
	switch os {
	case "linux":
		return linuxAutocomplete(text)
	case "darwin":
		return text, nil
	case "windows":
		return text, nil
	default:
		return text, nil
	}
}

func linuxAutocomplete(text string) (string, error) {

	return text, nil
}
