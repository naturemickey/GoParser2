package lex

import "fmt"

type lexerInner struct {
	currentState    *lexerState
	lastFinishState *lexerState
	from            cursor
	scanner         *scanner
	nfa             *nfa
	nextToken       *Token
}

//func (this *lexerInner) LA() *Token {
//	res := this.nextToken
//	if res == nil {
//		this.nextToken = this._nextToken()
//		return this.nextToken
//	}
//	return res
//}
//
//func (this *lexerInner) Pop() bool {
//	if this.nextToken == nil {
//		return false
//	}
//	this.nextToken = nil
//	return true
//}

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

func NewLexerInnerWithFile(filepath string) *lexerInner {
	return NewLexerInnerWithFileInner(filepath, NFA)
}

func NewLexerInnerWithFileInner(filepath string, nfa *nfa) *lexerInner {
	currentState := newLexerState()
	lastFinishState := (*lexerState)(nil)
	from := newCursor()
	scanner := NewScannerFromFile(filepath)
	return &lexerInner{
		currentState:    currentState,
		lastFinishState: lastFinishState,
		from:            from,
		scanner:         scanner,
		nfa:             nfa}
}

func NewLexerInnerWithCode(code string) *lexerInner {
	return NewLexerInnerWithCodeInner(code, NFA)
}

func NewLexerInnerWithCodeInner(code string, nfa *nfa) *lexerInner {
	currentState := newLexerState()
	lastFinishState := (*lexerState)(nil)
	from := newCursor()
	scanner := NewScannerFromCode(code)
	return &lexerInner{
		currentState:    currentState,
		lastFinishState: lastFinishState,
		from:            from,
		scanner:         scanner,
		nfa:             nfa}
}

func (this *lexerInner) NextToken() *Token {
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
				fmt.Printf("lexer_inner,error char '%s'， %s\n", string(char), this.scanner.cursor.string())
				// 报错之后忽略一个字母，继续前进
				continue
			} else {
				// 走不下去，但已经有了可以finish的状态。
				break
			}
		}

		if tokenType, ok := set.isFinish(); ok {
			this.lastFinishState = this.currentState.clone()
			if tokenType == GoLexerCOMMENT { // comment不要贪婪匹配
				break
			}
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
