package ast

import "GoParser2/lex"

type RecvStmt struct {
	// recvStmt: (expressionList ASSIGN | identifierList DECLARE_ASSIGN)? recvExpr = expression;
	expressionList *ExpressionList // todo 这里要研究一下等号后面是哪些类型的表达式（应该是有限制的）
	assign         *lex.Token
	identifierList *IdentifierList
	declare_assign *lex.Token
	recvExpr       *Expression
}

func VisitRecvStmt(lexer *lex.Lexer) *RecvStmt {
	clone := lexer.Clone()

	var identifierList = VisitIdentifierList(lexer)
	var expressionList *ExpressionList
	var assign *lex.Token
	var declare_assign *lex.Token

	clone1 := lexer.Clone()
	if identifierList == nil {
		expressionList = VisitExpressionList(lexer)
		if expressionList != nil {
			assign = lexer.LA()
			if assign.Type_() != lex.GoLexerASSIGN {
				assign = nil
				lexer.Recover(clone1)
			} else {
				lexer.Pop() // assign
			}
		}
	} else {
		declare_assign = lexer.LA()
		if declare_assign.Type_() != lex.GoLexerDECLARE_ASSIGN {
			declare_assign = nil
			lexer.Recover(clone1)
		} else {
			lexer.Pop() // declare_assign
		}
	}

	recvExpr := VisitExpression(lexer)
	if recvExpr == nil {
		lexer.Recover(clone)
		return nil
	}

	return &RecvStmt{expressionList: expressionList, assign: assign, identifierList: identifierList, declare_assign: declare_assign, recvExpr: recvExpr}
}
