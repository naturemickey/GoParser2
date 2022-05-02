package lex

import "fmt"

type lexer struct {
	currentState    *lexerState
	lastFinishState *lexerState
	from            cursor
	scanner         *scanner
	nfa             *nfa
}

type lexerState struct {
	cursor cursor
	states *stateSet
}

func (this *lexerState) clone() *lexerState {
	return &lexerState{cursor: this.cursor, states: this.states}
}

func newLexerState() *lexerState {
	return &lexerState{cursor: newCursor(), states: &stateSet{}}
}

func NewLexerWithFile(filepath string) *lexer {
	return NewLexerWithFileInner(filepath, NFA)
}

func NewLexerWithFileInner(filepath string, nfa *nfa) *lexer {
	currentState := newLexerState()
	lastFinishState := (*lexerState)(nil)
	fromState := newCursor()
	scanner := NewScannerFromFile(filepath)
	return &lexer{currentState, lastFinishState, fromState, scanner, nfa}
}

func NewLexerWithCode(code string) *lexer {
	return NewLexerWithCodeInner(code, NFA)
}

func NewLexerWithCodeInner(code string, nfa *nfa) *lexer {
	currentState := newLexerState()
	lastFinishState := (*lexerState)(nil)
	fromState := newCursor()
	scanner := NewScannerFromCode(code)
	return &lexer{currentState, lastFinishState, fromState, scanner, nfa}
}

func (this *lexer) NextToken() *token {
	if this.lastFinishState != nil {
		// 从上一次的结束开始
		this.from = this.lastFinishState.cursor
		this.currentState.cursor = this.lastFinishState.cursor
		this.scanner.cursor = this.lastFinishState.cursor
	}
	// 状态初始化
	this.currentState.states = this.nfa.start.eqSet()
	// 清空结束状态
	this.lastFinishState = nil

	for this.scanner.Next() {
		char := this.scanner.LA()
		set := this.currentState.states.accept(char)

		// 每向前一步，更新当前状态
		this.currentState.cursor = this.scanner.cursor
		this.currentState.states = set

		if set.isEmpty() {
			if this.lastFinishState == nil {
				fmt.Printf("error char '%s'， %s\n", string(char), this.scanner.cursor.string())
				// 报错之后忽略一个字母，继续前进
				continue
			} else {
				// 走不下去，但已经有了可以finish的状态。
				break
			}
		}

		if _, ok := set.isFinish(); ok {
			this.lastFinishState = this.currentState.clone()
		}
	}

	if this.lastFinishState != nil {
		literal := this.scanner.getTokenLiteral(this.from, this.lastFinishState.cursor)
		type_, _ := this.lastFinishState.states.isFinish()
		return NewToken(type_, literal, this.lastFinishState.cursor.line, this.lastFinishState.cursor.column)
	} else {
		return nil
	}
}
