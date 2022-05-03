package lex

type Lexer struct {
	tokens []*token
	index  int
}

func NewLexer(lexeri *lexerInner) *Lexer {
	var tokens []*token
	if token := lexeri._nextToken(); token != nil {
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

func (this *Lexer) LA() *token {
	return this.tokens[this.index]
}

func (this *Lexer) Pop() *token {
	res := this.LA()
	this.index++
	return res
}
