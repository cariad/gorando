package gorando

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	"strings"
	"testing"
	"unicode/utf8"
)

var validPools = []string{
	"a",
	"ab",
	"abc",
	"☺☻☹",
	"日a本b語ç",
	"ça",
}

var invalidStates = []struct {
	name            string
	pool            string
	count           int
	expectedMessage string
}{
	{"EmptyPool", "", 1, "pool is empty"},
	{"ZeroCount", "a", 0, "count (0) must be greater than zero"},
	{"NegCount", "a", -1, "count (-1) must be greater than zero"},
}

func TestGetRunesWithValidPools(t *testing.T) {
	for _, pool := range validPools {
		name := fmt.Sprintf("ValidPool_%s", pool)
		t.Run(name, func(t *testing.T) {
			expectAllRunesInPool(pool, t)
		})
	}
}

func TestGetRunesWithInvalidStates(t *testing.T) {
	for _, state := range invalidStates {
		name := fmt.Sprintf("InvalidState_%s", state.name)
		t.Run(name, func(t *testing.T) {
			testInvalidState(state.pool, state.count, state.expectedMessage, t)
		})
	}
}

func testInvalidState(pool string, count int, expectMsg string, t *testing.T) {
	_, err := GetRunes(pool, count)
	expectError(err, expectMsg, t)
}

func expectError(err error, expected string, t *testing.T) {
	if err == nil {
		t.Errorf("expected error \"%s\" but none returned", expected)
	}

	actual := err.Error()

	if actual != expected {
		t.Errorf("expected error \"%s\" but got \"%s\"", expected, actual)
	}
}

func TestGetRandomWhenUnableToGetInteger(t *testing.T) {
	const expect = "not implemented"

	intGetter := func(rand io.Reader, max *big.Int) (n *big.Int, err error) {
		return nil, errors.New(expect)
	}

	_, err := getRunes("abc", 1, intGetter)
	expectError(err, expect, t)
}

func expectAllRunesInPool(pool string, t *testing.T) {
	remaining := pool
	remainingTries := utf8.RuneCountInString(pool) * 100

	for utf8.RuneCountInString(remaining) > 0 && remainingTries > 0 {
		runes, err := GetRunes(pool, 1)

		if err != nil {
			t.Error(err)
		}

		rs := string(runes[0])

		if strings.Index(pool, rs) < 0 {
			t.Errorf("rune \"%s\" was unexpectedly returned from pool \"%s\"", rs, pool)
		}

		remainChunks := strings.SplitN(remaining, rs, 2)
		remaining = strings.Join(remainChunks, "")
		remainingTries--
	}

	if utf8.RuneCountInString(remaining) > 0 {
		t.Errorf("never returned \"%s\" from pool \"%s\"", remaining, pool)
	}
}
