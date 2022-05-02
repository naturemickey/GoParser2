package lex

type stateSet struct {
	set []*state
}

func (this *stateSet) addState(state *state) bool {
	for _, s := range this.set {
		if s.id == state.id {
			return false
		}
	}
	this.set = append(this.set, state)
	return true
}

func (this *stateSet) merge(that *stateSet) {
	for _, s := range that.set {
		this.addState(s)
	}
}

func (this *stateSet) accept(char rune) *stateSet {
	states := &stateSet{}

	for _, s := range this.set {
		states.merge(s.accept(char))
	}

	return states.eqSet()
}
func (this *stateSet) eqSet() *stateSet {
	states := &stateSet{}

	for _, s := range this.set {
		states.merge(s.eqSet())
	}

	return states
}

func (this *stateSet) isFinish() (TokenType, bool) {
	resType := GoLexerNone
	for _, s := range this.set {
		if s.type_ != GoLexerNone {
			if resType == GoLexerNone || s.type_ < resType {
				resType = s.type_
			}
		}
	}
	return resType, resType != GoLexerNone
}

func (this *stateSet) isEmpty() bool {
	return len(this.set) == 0
}
