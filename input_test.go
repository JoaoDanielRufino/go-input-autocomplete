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

func TestAutocompleteWithEmptyString(t *testing.T) {
	input := NewInput("test: ")

	currentTextResult := input.Autocomplete()

	if currentTextResult != nil {
		t.Errorf("Autocomplete with empty string should return nil, instead returned %v", currentTextResult)
	}
}

func TestAutocompleteOnNonLinuxOS(t *testing.T) {
	expectedAutocomplete := "a"

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		t.Skipf("Skip test because OS is %v", runtime.GOOS)
	}
	input := NewInput("test: ")
	input.AddChar('a')
	err := input.Autocomplete()

	if err != nil {
		t.Errorf("Autocomplete should not return error on windows machines, insted got error %v", err)
	}
	if input.currentText != expectedAutocomplete {
		t.Errorf("Autocomplete with amy string on windows should return the same string - %v instead returned %v", expectedAutocomplete, input.currentText)
	}
}

func TestAutocompleteOnLinuxOS(t *testing.T) {
	expectedAutocomplete := "/bin/bash"

	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		t.Skipf("Skip test because OS is %v", runtime.GOOS)
	}

	input := NewInput("test: ")
	for _, ch := range "/bin/bas" {
		input.AddChar(ch)
	}

	err := input.Autocomplete()

	if err != nil {
		t.Errorf("Autocomplete should not return error, insted got error %v", err)
	}

	AssertTextAndPosition(t, input, expectedAutocomplete, 9)
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
