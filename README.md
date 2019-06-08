# gorando

[![CircleCI](https://circleci.com/gh/cariad/gorando.svg?style=svg)](https://circleci.com/gh/cariad/gorando)

A golang package for dealing with randomness.

There are many packages like it, but this one is mine.

## Examples

```go
package main

import (
    "fmt"
    "github.com/cariad/gorando"
)

func main() {
    abc, _ := gorando.GetString("abc", 8)
    fmt.Println(abc)
    // bcbbcaab

    japanese, _ := gorando.GetString("こんにちは", 8)
    fmt.Println(japanese)
    // こにこちにんんは

    emoji, _ := gorando.GetString("🌈😎🐄", 8)
    fmt.Println(emoji)
    // 😎🐄😎🐄🐄🐄😎🌈
}
```

## Functions

```go
func GetRunes(pool string, count int) ([]rune, error)
```

`GetRunes` returns a slice of runes of length `count` populated by a random pick of runes out of the `pool` string.

```go
func GetString(pool string, length int) (string, error)
```

`GetString` returns a string of length `length`  populated by a random pick of runes out of the `pool` string.

## Security

This package uses `crypto/rand` to avoid the predictability of `math/random`.

Also, this package intentionally does not log any pools or random selections.
