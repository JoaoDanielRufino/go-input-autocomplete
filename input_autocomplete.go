package input_autocomplete

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func keyboardListener(input *Input) error {
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			return err
		}

		switch key {
		case keyboard.KeyEnter:
			fmt.Println("")
			return nil
		case keyboard.KeyArrowLeft:
			input.MoveCursorLeft()
		case keyboard.KeyArrowRight:
			input.MoveCursorRight()
		case keyboard.KeyBackspace:
			input.RemoveChar()
		case keyboard.KeyBackspace2:
			input.RemoveChar()
		case keyboard.KeyTab:
			err := input.Autocomplete()
			if err != nil {
				return err
			}

		default:
			input.AddChar(char)
		}
	}
}

func Read(text string) (string, error) {
	if err := keyboard.Open(); err != nil {
		return "", err
	}

	defer keyboard.Close()

	input := NewInput(text)

	input.Print()

	if err := keyboardListener(input); err != nil {
		return "", err
	}

	return input.GetCurrentText(), nil
}

func isDir(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return info.IsDir()
}