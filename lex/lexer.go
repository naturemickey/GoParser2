package lex

type Lexer struct {
	tokens []*Token
	index  int
}

func NewLexer(lexeri *lexerInner) *Lexer {
	var tokens []*Token
	for token := lexeri.NextToken(); token != nil; token = lexeri.NextToken() {
		if token.Type_() == GoLexerCOMMENT || token.Type_() == GoLexerLINE_COMMENT ||
			token.Type_() == GoLexerWS || token.Type_() == GoLexerTERMINATOR {
			continue
		}
		tokens = append(tokens, token)
	}
	return &Lexer{tokens: tokens, index: 0}
}

func NewLexerWithFile(filepath string) *Lexer {
	return NewLexer(NewLexerInnerWithFileInner(filepath, NFA))
}

func NewLexerWithCode(code string) *Lexer {
	return NewLexer(NewLexerInnerWithCodeInner(code, NFA))
}

func (this *Lexer) Clone() *Lexer {
	return &Lexer{tokens: this.tokens, index: this.index}
}

func (this *Lexer) Recover(that *Lexer) {
	this.tokens = that.tokens
	this.index = that.index
}

func (this *Lexer) LA() *Token {
	return this._la(0)
}

func (this *Lexer) LA0() *Token {
	return this._la(0)
}

func (this *Lexer) LA1() *Token {
	return this._la(1)
}

func (this *Lexer) LA2() *Token {
	return this._la(2)
}

func (this *Lexer) LA3() *Token {
	return this._la(3)
}

func (this *Lexer) _la(i int) *Token {
	if this.index+i < len(this.tokens) {
		return this.tokens[this.index+i]
	}
	return nil
}

func (this *Lexer) Pop() bool {
	if this.index < len(this.tokens) {
		this.index++
		return true
	}
	return false
}
