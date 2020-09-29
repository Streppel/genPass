<h1 align="center">
  <br>
  🔒 genPass 🔒
</h1>

Simple and straight-forward random password generation, offering simple but useful options for how the final password will look like.

This was born out of the need of a random string password generator to comply 

Currently we support the following usecases:
- Numeric passwords
- Alphabetic only passwords
- Alphanumeric passwords
- Alphanumeric with symbols (~!@#$%^&*()_-+?)

Please note that currently we do not support fancy stuff like entropy control at the moment, leaving it entirely in control of [crypto/rand](https://godoc.org/crypto/rand).

Examples:

```go
import (
	"fmt"

	"github.com/streppel/genpass"
)

func main() {
	// if unspecified in the parameters, default password length is 8
	// default password is numeric
	numericPassword := genpass.NewPassword()

	fmt.Println(numericPassword) // example: "34997406"

	// to generate an alphabetic password, use Alphabetic option
	alphabeticPassword := genpass.NewPassword(genpass.WithCharacters(genpass.Alphabetic))

	fmt.Println(alphabeticPassword) // example: "pduhkuyu"

	// so on with alphanumeric, or with symbols
	alphanumericPassword := genpass.NewPassword(genpass.WithCharacters(genpass.Alphanumeric))
	symbolsPassword := genpass.NewPassword(genpass.WithCharacters(genpass.AlphanumericWithSymbols))

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