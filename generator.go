package genpass

import (
	"crypto/rand"
	"io"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	symbols      = "~!@#$%^&*()_-+?"
)

type CharacterType int
type TypeCase int

type generator struct {
	CharacterType CharacterType
	TypeCase      TypeCase // only used when alpha characters are involved

	Length int

	randomnessGenerator io.Reader
}

func newGenerator(opts ...Param) *generator {
	g := &generator{
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

func (p generator) Generate() string {
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

func (p generator) getCharacter(b byte) rune {
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

func (p generator) digit(i int) rune {
	return p.getRuneFrom(digits, i)
}

func (p generator) alpha(i int) rune {
	return p.getRuneFrom(p.casedLetters(), i)
}

func (p generator) alphanum(i int) rune {
	return p.getRuneFrom(p.casedLetters()+digits, i)
}

func (p generator) alphanumWithSymbols(i int) rune {
	return p.getRuneFrom(p.casedLetters()+digits+symbols, i)
}

func (p generator) casedLetters() string {
	switch p.TypeCase {
	case Uppercase:
		return upperLetters
	case Lowercase:
		return lowerLetters
	default:
		return upperLetters + lowerLetters
	}
}

func (p generator) getRuneFrom(str string, i int) rune {
	return rune(str[i%len(str)])
}
