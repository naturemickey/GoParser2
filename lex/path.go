package lex

type path interface {
	accept(char rune) (*state, bool)
}

const e rune = 0

type charPath struct {
	char rune
	to   *state
}

func (this charPath) accept(char rune) (*state, bool) {
	if this.char == char {
		return this.to, true
	}
	return nil, false
}

type regPath struct {
	regular func(rune) bool
	to      *state
}

func (this regPath) accept(char rune) (*state, bool) {
	if this.regular(char) {
		return this.to, true
	}
	return nil, false
}

func pathEq(this, that path) bool {
	if p1, ok := this.(*charPath); ok {
		if p2, ok := that.(*charPath); ok {
			return p1.char == p2.char && p1.to == p2.to
		}
	}
	return false
}

var _ path = (*charPath)(nil)
var _ path = (*regPath)(nil)
