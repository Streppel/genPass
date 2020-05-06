package fluentpass

import (
	"github.com/streppel/fluentpass/password"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigCalled(t *testing.T) {
	var called = false
	password.NewGenerator(func(*password.Generator) {
		called = true
	})
	assert.True(t, called)
}

func TestConfigurationSet(t *testing.T) {
	t.Run(`test WithCharacters changes characterType state`, func(t *testing.T) {
		assert.Equal(t, password.NewGenerator(WithCharacters(password.Alphabetic)).CharacterType, password.Alphabetic)
		assert.Equal(t, password.NewGenerator(WithCharacters(password.Alphanumeric)).CharacterType, password.Alphanumeric)
		assert.Equal(t, password.NewGenerator(WithCharacters(password.Numeric)).CharacterType, password.Numeric)
	})

	t.Run(`test WithLength changes length state`, func(t *testing.T) {
		assert.Equal(t, password.NewGenerator(WithLength(5)).Length, 5)
		assert.Equal(t, password.NewGenerator(WithLength(10)).Length, 10)
	})

	t.Run(`test WithLength handles negative numbers`, func(t *testing.T) {
		assert.Equal(t, password.NewGenerator(WithLength(-5)).Length, 0)
	})
}
