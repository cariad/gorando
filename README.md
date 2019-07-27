# gorando

[![Build Status](https://travis-ci.org/cariad/gorando.svg?branch=master)](https://travis-ci.org/cariad/gorando) [![Go Report Card](https://goreportcard.com/badge/github.com/cariad/gorando)](https://goreportcard.com/report/github.com/cariad/gorando) [![](https://godoc.org/github.com/cariad/gorando?status.svg)](http://godoc.org/github.com/cariad/gorando) [![MIT](https://img.shields.io/npm/l/express.svg)](https://github.com/cariad/gorando/blob/master/LICENSE)

A Golang package for dealing with randomness.

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
}
```

The pool can include non-English characters, including emoji:

```go
package main

import (
    "fmt"
    "github.com/cariad/gorando"
)

func main() {
    japanese, _ := gorando.GetString("こんにちは", 8)
    fmt.Println(japanese)
    // こにこちにんんは

    emoji, _ := gorando.GetString("🌈😎🐄", 8)
    fmt.Println(emoji)
    // 😎🐄😎🐄🐄🐄😎🌈
}
```

For quick access to common types of pool, there are `LowerAlpha`, `UpperAlpha` and `Digits` constants:

```go
package main

import (
    "fmt"
    "github.com/cariad/gorando"
)

func main() {
    onlyLowerAlpha, _ := gorando.GetString(gorando.LowerAlpha, 8)
    fmt.Println(onlyLowerAlpha)
    // lnosdvql

    onlyAlpha, _ := gorando.GetString(gorando.LowerAlpha + gorando.UpperAlpha, 8)
    fmt.Println(onlyAlpha)
    // uvYJfgwU

    onlyAlphaNum, _ := gorando.GetString(gorando.LowerAlpha + gorando.UpperAlpha + gorando.Digits, 8)
    fmt.Println(onlyAlphaNum)
    // 4Pz453II
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

## Constants

`LowerAlpha` is a string containing all the characters in the English alphabet in lower-case form.

`UpperAlpha` is a string containing all the characters in the English alphabet in upper-case form.

`Digits` is a string containing all the base-10 digits.

## Security

This package uses `crypto/rand` to avoid the predictability of `math/random`.

Also, this package intentionally does not log any pools or random selections.
