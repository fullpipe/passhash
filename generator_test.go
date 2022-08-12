package passhash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenerator(t *testing.T) {
	g := NewGenerator()

	g.
		Length(10).
		NumDigits(1).
		NumSymbols(2).
		NoUpper().
		AllowRepeat()

	assert.Equal(t, 10, g.length)
	assert.Equal(t, 1, g.numDigits)
	assert.Equal(t, 2, g.numSymbols)
	assert.True(t, g.noUpper)
	assert.True(t, g.allowRepeat)

	g.WithUpper().DisallowRepeat()

	assert.False(t, g.noUpper)
	assert.False(t, g.allowRepeat)

	p, err := g.Generate()

	assert.Len(t, p, 10)
	assert.NoError(t, err)
}
