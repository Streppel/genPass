package password

import (
	"crypto/rand"
	"io"
)

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

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	symbols      = "~!@#$%^&*()_-+?"
)

// public builder interface
type Param func(*Generator)

type Generator struct {
	CharacterType CharacterType
	TypeCase      TypeCase // only used when alpha characters are involved

	Length int

	randomnessGenerator io.Reader
}

func NewGenerator(opts ...Param) *Generator {
	g := &Generator{
		Length:              8,
		CharacterType:       Numeric,
		TypeCase:            Lowercase,
		randomnessGenerator: rand.Reader,
	}
	for _, f := range opts {
		f(g)
	}
	return g
}

func (p Generator) Generate() string {
	pwdLength := p.Length
	randValues := make([]byte, pwdLength)
	_, err := p.randomnessGenerator.Read(randValues)
	if err != nil {
		return ""
	}

	pwd := make([]rune, len(randValues))
	for i, randVal := range randValues {
		pwd[i] = p.getCharacter(randVal)
	}

	return string(pwd)
}

func (p Generator) getCharacter(b byte) rune {
	i := int(b)
	switch p.CharacterType {
	case AlphanumericWithSymbols:
		return p.alphanumWithSymbols(i)
	case Alphabetic:
		return p.alpha(i)
	case Alphanumeric:
		return p.alphanum(i)
	default:
		return p.digit(i)
	}
}

func (p Generator) digit(i int) rune {
	return p.getRuneFrom(digits, i)
}

func (p Generator) alpha(i int) rune {
	return p.getRuneFrom(p.casedLetters(), i)
}

func (p Generator) alphanum(i int) rune {
	return p.getRuneFrom(p.casedLetters()+digits, i)
}

func (p Generator) alphanumWithSymbols(i int) rune {
	return p.getRuneFrom(p.casedLetters()+digits+symbols, i)
}

func (p Generator) casedLetters() string {
	switch p.TypeCase {
	case Uppercase:
		return upperLetters
	case Lowercase:
		return lowerLetters
	default:
		return upperLetters + lowerLetters
	}
}

func (p Generator) getRuneFrom(str string, i int) rune {
	return rune(str[i%len(str)])
}
