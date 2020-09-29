<h1 align="center">
  <br>
  🔢 genPass 🔣
</h1>

Simple and straight-forward **random password generation** in Go, offering simple but useful options for how the final password will look like.

Currently we support the following usecases:
- Numeric passwords
- Alphabetic passwords
- Alphanumeric passwords
- Alphanumeric with symbols (~!@#$%^&*()_-+?)

Sample of generated passwords (using default length with various charset and case options options):
```
93376358    xnzrowlk    34weaqb0
0oc7lm7?    146EQGAQ    ye7OjZmz
```

Please note that currently we do not support "fancy stuff" like entropy control at the moment, leaving it entirely in control of Go's [crypto/rand](https://godoc.org/crypto/rand) package.

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
	alphabeticPass := genpass.NewPassword(genpass.WithCharacters(genpass.Alphabetic))
	fmt.Println(alphabeticPass) // example: "pduhkuyu"

	// so on with alphanumeric, or with symbols
	alphanumericPass := genpass.NewPassword(
		genpass.WithCharacters(genpass.Alphanumeric))
	symbolsPass := genpass.NewPassword(
		genpass.WithCharacters(genpass.AlphanumericWithSymbols))
	fmt.Println(alphanumericPass) // example: "7bvimig3"
	fmt.Println(symbolsPass) // example: "~sa&&*_c"

	// when using alpha characters, we can also modify the type case (default is lower)
	alphaUpperPass := genpass.NewPassword(
		genpass.WithCharacters(genpass.Alphanumeric),
		genpass.WithCase(genpass.Uppercase))
	alphanumericMixedPass := genpass.NewPassword(
		genpass.WithCharacters(genpass.Alphanumeric),
		genpass.WithCase(genpass.Mixedcase))
	fmt.Println(alphaUpperPass) // example: "I4U1ZHIV"
	fmt.Println(alphanumericMixedPass) // example: "6YTNzu2w"
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
