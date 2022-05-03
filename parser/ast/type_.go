package ast

import "GoParser2/lex"

type Type_ struct {
}

func VisitType_(lexer *lex.Lexer) *Type_ {
	// 识别失败不是真失败，需要恢复lexer，因为外面可能会继续使用
	// clone := lexer.Clone()
	panic("todo")
}
