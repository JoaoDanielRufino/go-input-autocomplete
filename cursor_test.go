package input_autocomplete

import "testing"

func TestNewCursor(t *testing.T) {
	c := NewCursor()
	expected := Cursor{position: 0}
	if *c != expected {
		t.Errorf("Expected %v", expected)
	}
}

func TestIncrementPosition(t *testing.T) {
	c := NewCursor()
	expected := Cursor{position: 1}
	c.IncrementPosition()
	if *c != expected {
		t.Errorf("Expected %v", expected)
	}
}

func TestSetPosition(t *testing.T) {
	c := NewCursor()
	expected := Cursor{position: 69}
	c.SetPosition(69)
	if *c != expected {
		t.Errorf("Expected %v", expected)
	}
}

func TestMoveRight(t *testing.T) {
	c := NewCursor()
	expected := Cursor{position: 1}
	c.MoveRight()
	if *c != expected {
		t.Errorf("Expected %v", expected)
	}
}

func TestMoveLeft(t *testing.T) {
	c := &Cursor{position: 421}
	expected := Cursor{position: 420}
	c.MoveLeft()
	if *c != expected {
		t.Errorf("Expected %v", expected)
	}
}

func TestMoveLeftNPos(t *testing.T) {
	c := &Cursor{position: 50}
	expected := Cursor{position: 25}
	c.MoveLeftNPos(25)
	if *c != expected {
		t.Errorf("Expected %v", expected)
	}
}

func TestGetPosition(t *testing.T) {
	c := &Cursor{position: 20}
	if c.GetPosition() != 20 {
		t.Errorf("Expected 20")
	}
}
