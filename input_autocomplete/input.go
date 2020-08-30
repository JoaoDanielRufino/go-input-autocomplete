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

func (i *Input) AddChar(c string) {
	i.currentText += c
	i.cursor.IncrementPosition()
}

func (i *Input) canDeleteChar() bool {
	return len(i.currentText) >= 1
}

func (i *Input) RemoveChar() {
	if i.canDeleteChar() {
		pos := i.cursor.GetPosition()
		i.currentText = i.currentText[:pos-1] + i.currentText[pos:]
		i.cursor.SetPosition(len(i.currentText))
	}
}

func (i *Input) MoveCursorLeft() {
	i.cursor.MoveLeft()
}

func (i *Input) MoveCursorLeftTo(x int) {
	i.cursor.MoveLeftTo(x)
}

func (i *Input) MoveCursorRight() {
	i.cursor.MoveRight()
}

func (i *Input) Print() {
	fmt.Print("\033[G\033[K")
	fmt.Print(i.fixedText + i.currentText)
}

func (i *Input) GetCurrentText() string {
	return i.currentText
}
