package gorando

import (
	"testing"
	"unicode/utf8"
)

func TestGetStringFailsForZeroLength(t *testing.T) {
	const expected = "count (0) must be greater than zero"
	_, err := GetString("a", 0)

	if err == nil {
		t.Errorf("expected error \"%s\" but none returned", expected)
	}

	actual := err.Error()

	if actual != expected {
		t.Errorf("expected error \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestGetStringReturnsCorrectLength(t *testing.T) {
	const maxLength = 128 // Nothing special; just a large length to test.
	for length := 1; length <= maxLength; length++ {
		testStringLength(length, t)
	}
}

func testStringLength(length int, t *testing.T) {
	s, err := GetString("a", length)

	if err != nil {
		t.Error(err)
	}

	actual := utf8.RuneCountInString(s)

	if actual != length {
		t.Errorf("expected %d runes but got %d (%s)", length, actual, s)
	}
}
