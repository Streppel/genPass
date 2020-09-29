package fluentpass

import (
	"testing"

	"github.com/streppel/genpass/internal"

	"github.com/stretchr/testify/assert"
)

func TestConfigCalled(t *testing.T) {
	var called = false
	internal.NewGenerator(func(generator *internal.Generator) {
		called = true
	})
	assert.True(t, called)
}

func TestConfigurationSet(t *testing.T) {
	t.Run(`test WithCharacters changes characterType state`, func(t *testing.T) {
		assert.Equal(t, internal.NewGenerator(WithCharacters(Alphabetic)).CharacterType, Alphabetic)
		assert.Equal(t, internal.NewGenerator(WithCharacters(Alphanumeric)).CharacterType, Alphanumeric)
		assert.Equal(t, internal.NewGenerator(WithCharacters(Numeric)).CharacterType, Numeric)
	})

	t.Run(`test WithLength changes length state`, func(t *testing.T) {
		assert.Equal(t, internal.NewGenerator(WithLength(5)).Length, 5)
		assert.Equal(t, internal.NewGenerator(WithLength(10)).Length, 10)
	})

	t.Run(`test WithLength handles negative numbers`, func(t *testing.T) {
		assert.Equal(t, internal.NewGenerator(WithLength(-5)).Length, 0)
	})
}
