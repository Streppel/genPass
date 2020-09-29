package genpass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigCalled(t *testing.T) {
	var called = false
	newGenerator(func(generator *generator) {
		called = true
	})
	assert.True(t, called)
}

func TestConfigurationSet(t *testing.T) {
	t.Run(`test WithCharacters changes characterType state`, func(t *testing.T) {
		assert.Equal(t, newGenerator(WithCharacters(Alphabetic)).CharacterType, Alphabetic)
		assert.Equal(t, newGenerator(WithCharacters(Alphanumeric)).CharacterType, Alphanumeric)
		assert.Equal(t, newGenerator(WithCharacters(Numeric)).CharacterType, Numeric)
	})

	t.Run(`test WithLength changes length state`, func(t *testing.T) {
		assert.Equal(t, newGenerator(WithLength(5)).Length, 5)
		assert.Equal(t, newGenerator(WithLength(10)).Length, 10)
	})

	t.Run(`test WithLength handles negative numbers`, func(t *testing.T) {
		assert.Equal(t, newGenerator(WithLength(-5)).Length, 0)
	})
}
