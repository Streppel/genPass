<h1 align="center">
  <br>
  🔢 genPass 🔣
</h1>

Simple and straight-forward **random password generation**, offering simple but useful options for how the final password will look like.

This was born out of the need of a... random string password generator, with the bonus of choosing what kind of characters will be in it.

Currently we support the following usecases:
- Numeric passwords
- Alphabetic only passwords
- Alphanumeric passwords
- Alphanumeric with symbols (~!@#$%^&*()_-+?)

Please note that currently we do not support "fancy stuff" like entropy control at the moment, leaving it entirely in control of [crypto/rand](https://godoc.org/crypto/rand).

Examples:

```go
import (
	"fmt"

	"github.com/streppel/genpass"
)

func main() {
	// default password length is 8
	// default password is numeric
	numericPass := genpass.NewPassword()
	fmt.Println(numericPass) // example: "34997406"

	// control password len with .WithLength
	tinyPass := genpass.NewPassword(genpass.WithLength(3))
	fmt.Println(tinyPass) // example: "531"

	// to generate an alphabetic password, use Alphabetic option
	alphabeticPassword := genpass.NewPassword(
		genpass.WithCharacters(genpass.Alphabetic))
	fmt.Println(alphabeticPassword) // example: "pduhkuyu"

	// so on with alphanumeric, or with symbols
	alphanumericPassword := genpass.NewPassword(
		genpass.WithCharacters(genpass.Alphanumeric))
	symbolsPassword := genpass.NewPassword(
		genpass.WithCharacters(genpass.AlphanumericWithSymbols))
	fmt.Println(alphanumericPassword) // example: "7bvimig3"
	fmt.Println(symbolsPassword) // example: "~sa&&*_c"

	// when using alpha characters, we can also modify the type case (default is lower)
	alphaUpperPassword := genpass.NewPassword(
		genpass.WithCharacters(genpass.Alphanumeric),
		genpass.WithCase(genpass.Uppercase))
	alphanumericMixedPassword := genpass.NewPassword(
		genpass.WithCharacters(genpass.Alphanumeric),
		genpass.WithCase(genpass.Mixedcase))
	fmt.Println(alphaUpperPassword) // example: "I4U1ZHIV"
	fmt.Println(alphanumericMixedPassword) // example: "6YTNzu2w"
}
```

Benchmarking with default password generation:
```
goos: linux
goarch: amd64
pkg: github.com/streppel/genpass
BenchmarkNewPassword-8           1819519               653 ns/op
PASS
```
