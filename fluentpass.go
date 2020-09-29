package genpass

type Param func(generator *generator)

// NewPassword returns a string representing password. It returns an empty string in case of error.
// It accepts functional parameters, such as password length, character type and case-sensitiviness.
// Examples:
// NewPassword() // returns a numeric password with default length (length 8)
// NewPassword(WithLength(12)) // returns a numeric password with length 12
// NewPassword(WithCharacters(Alphanumeric), WithLength(12)) // returns an alphamumeric password with length 12
func NewPassword(opts ...Param) string {
	return newGenerator(opts...).Generate()
}

func WithCharacters(t characterType) Param {
	return func(generator *generator) {
		generator.CharacterType = t
	}
}

func WithLength(i int) Param {
	return func(generator *generator) {
		if i < 0 {
			i = 0
		}
		generator.Length = i
	}
}

const (
	Numeric characterType = iota
	Alphabetic
	Alphanumeric
	AlphanumericWithSymbols
)

const (
	Lowercase typeCase = iota
	Uppercase
	Mixedcase
)
