package lex

import "sync/atomic"

type state struct {
	id    uint32
	type_ TokenType
	paths []path
}

var stateIdSequence uint32 = 0

func nextStateId() uint32 {
	return atomic.AddUint32(&stateIdSequence, 1)
}

func (this *state) accept(char rune) *stateSet {
	states := &stateSet{}

	for _, p := range this.paths {
		if s, ok := p.accept(char); ok {
			// states.merge(s.eqSet())
			states.addState(s)
		}
	}

	return states
}

func (this *state) eqSet() *stateSet {
	states := &stateSet{}
	states.addState(this)
	_eqSet(states, this)
	return states
}

func _eqSet(states *stateSet, this *state) {
	for _, p := range this.paths {
		if s, ok := p.accept(e); ok {
			if states.addState(s) {
				_eqSet(states, s)
			}
		}
	}
}

func (this *state) SetType(type_ TokenType) {
	this.type_ = type_
}

func NewState() *state {
	return &state{id: nextStateId()}
}

// 调用这个函数的必须是finish节点
func NewStateWithType(type_ TokenType) *state {
	if type_ == GoLexerNone {
		return NewState()
	}
	return &state{id: nextStateId(), type_: type_}
}

func (this *state) addCharPath(char rune, to *state) {
	path := &charPath{char, to}
	for _, p := range this.paths {
		if pathEq(p, path) {
			return
		}
	}
	this.paths = append(this.paths, path)
}

func (this *state) addRegularPath(regular func(rune) bool, to *state) {
	this.paths = append(this.paths, &regPath{regular, to})
}
