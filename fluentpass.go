package fluentpass

import "github.com/streppel/fluentpass/password"

// NewPassword returns a string representing password. It returns an empty string in case of error.
// It accepts functional parameters, such as password length, character type and case-sensitiviness.
// Examples:
// NewPassword() // returns a numeric password with default length (length 8)
// NewPassword(WithLength(12)) // returns a numeric password with length 12
// NewPassword(WithCharacters(Alphanumeric), WithLength(12)) // returns an alphamumeric password with length 12
func NewPassword(opts ...password.Param) string {
	p := password.NewGenerator(opts...)
	return p.Generate()
}

func WithCharacters(t password.CharacterType) password.Param {
	return func(generator *password.Generator) {
		generator.CharacterType = t
	}
}

func WithLength(i int) password.Param {
	return func(generator *password.Generator) {
		if i < 0 {
			i = 0
		}
		generator.Length = i
	}
}
