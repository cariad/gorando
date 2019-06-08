package gorando

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	"unicode/utf8"
)

// intGetter describes a function which can return a random integer. This is
// used only to test internal error handling, and allows tests to use a mocked
// "rand.Int" instead of the real thing.
type intGetter func(rand io.Reader, max *big.Int) (n *big.Int, err error)

// GetRunes returns a slice of runes selected randomly from the pool.
func GetRunes(pool string, count int) ([]rune, error) {
	return getRunes(pool, count, rand.Int)
}

func getRunes(pool string, count int, ig intGetter) ([]rune, error) {
	if count <= 0 {
		msg := fmt.Sprintf("count (%d) must be greater than zero", count)
		return make([]rune, 0), errors.New(msg)
	}

	poolLength := utf8.RuneCountInString(pool)
	resultRunes := make([]rune, count)

	if poolLength == 0 {
		return resultRunes, errors.New("pool is empty")
	}

	poolRunes := []rune(pool)
	maxPoolIndex := big.NewInt(int64(poolLength))

	for i := 0; i < count; i++ {
		v, err := ig(rand.Reader, maxPoolIndex)
		if err != nil {
			return resultRunes, err
		}
		poolIndex := v.Int64()
		resultRunes[i] = poolRunes[poolIndex]
	}

	return resultRunes, nil
}
