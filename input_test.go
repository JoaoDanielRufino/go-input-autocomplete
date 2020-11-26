package input_autocomplete

import (
	"runtime"
	"testing"
)

func TestCanDeleteCharWhenPositionIsStartOfText(t *testing.T) {
	// Position is in the start of text
	input := NewInput("test: ")
	input.AddChar('1')
	input.cursor.SetPosition(0)

	canDeleteCharResult := input.canDeleteChar()
	if canDeleteCharResult == true {
		t.Errorf("when position is 0 should enable deleting")
	}
}

func TestCanDeleteCharWhenPositionIsEndOfText(t *testing.T) {
	// Move position to end of text
	input := NewInput("test: ")
	input.AddChar('1')

	canDeleteCharResult := input.canDeleteChar()
	if canDeleteCharResult == false {
		t.Errorf("when position is in the end of the text should disable deleting")
	}
}

func TestAddCharInEndOfText(t *testing.T) {
	input := NewInput("test: ")
	input.AddChar('a')
	AssertTextAndPosition(t, input, "a", 1)
	input.AddChar('b')
	AssertTextAndPosition(t, input, "ab", 2)
}

func TestAddCharInMiddleOfText(t *testing.T) {
	input := NewInput("test: ")
	input.AddChar('a')
	AssertTextAndPosition(t, input, "a", 1)
	input.AddChar('b')
	AssertTextAndPosition(t, input, "ab", 2)
	input.cursor.MoveLeft()
	input.AddChar('c')
	AssertTextAndPosition(t, input, "acb", 2)
}

func TestRemoveChar(t *testing.T) {
	input := NewInput("test: ")

	input.AddChar('a')
	input.RemoveChar()

	AssertTextAndPosition(t, input, "", 0)
}

func TestRemoveCharWithOneLetter(t *testing.T) {
	input := NewInput("test: ")

	input.AddChar('a')
	input.RemoveChar()

	AssertTextAndPosition(t, input, "", 0)
}

func TestRemoveCharWhenInStartPosition(t *testing.T) {
	input := NewInput("test: ")

	input.AddChar('a')
	input.MoveCursorLeft()

	input.RemoveChar()

	// should not delete anything
	AssertTextAndPosition(t, input, "a", 0)
}

func TestRemoveCharWithThreeLetter(t *testing.T) {
	input := NewInput("test: ")

	input.AddChar('a')
	input.AddChar('b')
	input.AddChar('c')

	input.MoveCursorLeft()
	input.RemoveChar()

	AssertTextAndPosition(t, input, "ac", 1)
}

func TestMoveCursorLeft(t *testing.T) {
	input := NewInput("test: ")

	input.AddChar('a')
	input.MoveCursorLeft()

	AssertPosition(t, input, 0)
}

func TestMoveCursorRightWhenInStartPosition(t *testing.T) {
	input := NewInput("test: ")

	input.AddChar('a')
	input.cursor.SetPosition(0)
	input.MoveCursorRight()

	AssertPosition(t, input, 1)
}

func TestMoveCursorRightWhenInEndPosition(t *testing.T) {
	input := NewInput("test: ")

	input.AddChar('a')
	input.MoveCursorRight()

	AssertPosition(t, input, 1)
}

func TestGetCurrentText(t *testing.T) {
	expectedText := "a"
	input := NewInput("test: ")

	input.AddChar('a')
	currentTextResult := input.GetCurrentText()

	if currentTextResult != expectedText {
		t.Errorf("Current text should be equal to %v, got %v ", expectedText, currentTextResult)
	}
}

func TestAutocompleteOnNonUnixOS(t *testing.T) {
	expectedAutocomplete := "C:\\ProgramData\\"

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		t.Skipf("Skip test because OS is %v", runtime.GOOS)
	}
	input := NewInput("test: ")
	for _, ch := range "C:\\ProgramDa" {
		input.AddChar(ch)
	}

	input.Autocomplete()

	AssertTextAndPosition(t, input, expectedAutocomplete, len(expectedAutocomplete))
}

func TestAutocompleteOnUnixOS(t *testing.T) {
	expectedAutocomplete := "/etc/passwd"

	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		t.Skipf("Skip test because OS is %v", runtime.GOOS)
	}

	input := NewInput("test: ")
	for _, ch := range "/etc/passw" {
		input.AddChar(ch)
	}

	input.Autocomplete()

	AssertTextAndPosition(t, input, expectedAutocomplete, len(expectedAutocomplete))
}

func AssertTextAndPosition(t *testing.T, input *Input, expectedText string, expectedPosition int) {
	AssertText(t, input, expectedText)
	AssertPosition(t, input, expectedPosition)
}

func AssertText(t *testing.T, input *Input, expectedText string) {
	if input.currentText != expectedText {
		t.Errorf("Current text should be equal to %v, got %v ", expectedText, input.currentText)
	}
}

func AssertPosition(t *testing.T, input *Input, expectedPosition int) {
	if input.cursor.GetPosition() != expectedPosition {
		t.Errorf("current position should be equal to %v, got %v ", expectedPosition, input.cursor.GetPosition())
	}
}
