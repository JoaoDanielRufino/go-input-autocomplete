package input_autocomplete

import "fmt"

type Input struct {
	cursor      *Cursor
	fixedText   string
	currentText string
}

func NewInput(fixedText string) *Input {
	return &Input{
		cursor:      NewCursor(),
		fixedText:   fixedText,
		currentText: "",
	}
}

func (i *Input) canDeleteChar() bool {
	return i.cursor.GetPosition() >= 1
}

func (i *Input) AddChar(char rune) {
	pos := i.cursor.GetPosition()
	c := string(char)

	if pos == len(i.currentText) {
		i.currentText += c
		fmt.Print(c)
		i.cursor.IncrementPosition()
	} else {
		aux := len(i.currentText) - pos
		i.currentText = i.currentText[:pos] + c + i.currentText[pos:]
		i.cursor.SetPosition(len(i.currentText))
		i.Print()
		i.cursor.MoveLeftNPos(aux)
	}
}

func (i *Input) RemoveChar() {
	if i.canDeleteChar() {
		pos := i.cursor.GetPosition()
		aux := len(i.currentText) - pos
		i.currentText = i.currentText[:pos-1] + i.currentText[pos:]
		i.cursor.SetPosition(len(i.currentText))
		i.Print()
		i.cursor.MoveLeftNPos(aux)
	}
}

func (i *Input) MoveCursorLeft() {
	i.cursor.MoveLeft()
}

func (i *Input) MoveCursorRight() {
	if i.cursor.GetPosition() < len(i.currentText) {
		i.cursor.MoveRight()
	}
}

func (i *Input) Autocomplete() error {
	autocompletedText, err := Autocomplete(i.currentText)
	if err != nil {
		return err
	}

	i.currentText = autocompletedText
	i.cursor.SetPosition(len(i.currentText))
	i.Print()

	return err
}

func (i *Input) Print() {
	fmt.Print("\033[G\033[K")
	fmt.Print(i.fixedText + i.currentText)
}

func (i *Input) GetCurrentText() string {
	return i.currentText
}
