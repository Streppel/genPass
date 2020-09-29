package fluentpass

import (
	"github.com/streppel/genpass/internal"
)

type Param func(generator *internal.Generator)

// NewPassword returns a string representing password. It returns an empty string in case of error.
// It accepts functional parameters, such as password length, character type and case-sensitiviness.
// Examples:
// NewPassword() // returns a numeric password with default length (length 8)
// NewPassword(WithLength(12)) // returns a numeric password with length 12
// NewPassword(WithCharacters(Alphanumeric), WithLength(12)) // returns an alphamumeric password with length 12
func NewPassword(opts ...Param) string {
	p := internal.NewGenerator(opts...)
	return p.Generate()
}

func WithCharacters(t CharacterType) Param {
	return func(generator *internal.Generator) {
		generator.CharacterType = t
	}
}

func WithLength(i int) Param {
	return func(generator *internal.Generator) {
		if i < 0 {
			i = 0
		}
		generator.Length = i
	}
}

type CharacterType int

const (
	Numeric CharacterType = iota
	Alphabetic
	Alphanumeric
	AlphanumericWithSymbols
)

type TypeCase int

const (
	Lowercase TypeCase = iota
	Uppercase
	Mixedcase
)
