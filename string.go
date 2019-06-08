package gorando

import (
	"strings"
)

// GetString returns a random string constructed from runes in the pool.
func GetString(pool string, length int) (string, error) {
	runes, err := GetRunes(pool, length)

	if err != nil {
		return "", err
	}

	var sb strings.Builder

	for _, r := range runes {
		sb.WriteRune(r)
	}

	return sb.String(), err
}
