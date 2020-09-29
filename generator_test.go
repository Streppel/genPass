package genpass

import (
	"strings"
	"testing"
	"unicode"
)

type sequentialReader struct {
	// used only to fulfill io.Reader interface
}

func (s sequentialReader) Read(p []byte) (n int, err error) {
	for i, _ := range p {
		p[i] = byte(i)
	}
	return 0, nil
}

func TestPasswordContent(t *testing.T) {
	g := newGenerator()
	g.Length = 1 << 5

	g.randomnessGenerator = sequentialReader{}

	t.Run(`with digits only`, func(t *testing.T) {
		// build
		g.CharacterType = Numeric

		// execute
		pwd := g.Generate()

		// test
		for _, c := range pwd {
			if !unicode.IsDigit(c) {
				t.Errorf("pwd should only contain numeric characters\npwd:%s", pwd)
				t.FailNow()
			}
		}
	})

	t.Run(`with alphanumerical characters only`, func(t *testing.T) {
		// build
		g.CharacterType = Alphanumeric

		// execute
		pwd := g.Generate()

		// test
		for _, c := range pwd {
			if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
				t.Errorf("pwd should only contain alphanumeric characters\npwd:%s", pwd)
				t.FailNow()
			}
		}
	})

	t.Run(`with alphabetic characters only`, func(t *testing.T) {
		// build
		g.CharacterType = Alphabetic

		// execute
		pwd := g.Generate()

		// test
		for _, c := range pwd {
			if !unicode.IsLetter(c) {
				t.Errorf("pwd should only contain alphabetic characters\npwd:%s", pwd)
				t.FailNow()
			}
		}
	})

	t.Run(`with alphanumerical & special characters`, func(t *testing.T) {
		// build
		g.CharacterType = AlphanumericWithSymbols

		// execute
		pwd := g.Generate()

		// test
		for _, c := range pwd {
			if !unicode.IsDigit(c) && !unicode.IsLetter(c) && !strings.ContainsRune(symbols, c) {
				t.Errorf("pwd should only contain alphanumeric characters or allowed symbols\npwd:%s", pwd)
				t.FailNow()
			}
		}
	})

	t.Run(`with uppercase only`, func(t *testing.T) {
		// build
		g.CharacterType = Alphabetic
		g.TypeCase = Uppercase

		// execute
		pwd := g.Generate()

		// test
		for _, c := range pwd {
			if !unicode.IsLetter(c) && !unicode.IsUpper(c) {
				t.Errorf("pwd should only contain uppercase letters\npwd:%s", pwd)
				t.FailNow()
			}
		}
	})

	t.Run(`with lower only`, func(t *testing.T) {
		// build
		g.CharacterType = Alphabetic
		g.TypeCase = Lowercase

		// execute
		pwd := g.Generate()

		// test
		for _, c := range pwd {
			if !unicode.IsLetter(c) && !unicode.IsLower(c) {
				t.Errorf("pwd should only contain lowercase letters\npwd:%s", pwd)
				t.FailNow()
			}
		}
	})

	t.Run(`with mixed case`, func(t *testing.T) {
		// build
		g.CharacterType = Alphabetic
		g.TypeCase = Mixedcase

		// execute
		pwd := g.Generate()

		// test
		for _, c := range pwd {
			if !unicode.IsLower(c) && !unicode.IsUpper(c) {
				t.Errorf("pwd should only contain lowercase or uppercase letters\npwd:%s", pwd)
				t.FailNow()
			}
		}
	})
}
