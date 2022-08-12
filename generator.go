package passhash

import "github.com/sethvargo/go-password/password"

type Generator struct {
	length, numDigits, numSymbols int
	noUpper, allowRepeat          bool
}

func NewGenerator() *Generator {
	return &Generator{}
}

// Length set generated password length
func (g *Generator) Length(length int) *Generator {
	g.length = length

	return g
}

// NumDigits set number of digits in generated password
func (g *Generator) NumDigits(numDigits int) *Generator {
	g.numDigits = numDigits

	return g
}

// NumSymbols set number of symbols in generated password
func (g *Generator) NumSymbols(numSymbols int) *Generator {
	g.numSymbols = numSymbols

	return g
}

// NoUpper no uppercase in passwords
func (g *Generator) NoUpper() *Generator {
	g.noUpper = true

	return g
}

// WithUpper allow uppercase in passwords
func (g *Generator) WithUpper() *Generator {
	g.noUpper = false

	return g
}

// AllowRepeat allows repeats in passwords
func (g *Generator) AllowRepeat() *Generator {
	g.allowRepeat = true

	return g
}

// DisallowRepeat allows repeats in passwords
func (g *Generator) DisallowRepeat() *Generator {
	g.allowRepeat = false

	return g
}

// Generate generates some random password with rules
func (g *Generator) Generate() (string, error) {
	return password.Generate(g.length, g.numDigits, g.numSymbols, g.noUpper, g.allowRepeat)
}
