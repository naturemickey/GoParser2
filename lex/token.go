package lex

import "fmt"

type Token struct {
	type_   TokenType
	literal string
	line    int
	column  int
}

func (this *Token) Type_() TokenType {
	return this.type_
}

func (this *Token) Literal() string {
	return this.literal
}

func (this *Token) Line() int {
	return this.line
}

func (this *Token) Column() int {
	return this.column
}

func NewToken(type_ TokenType, literal string, line int, column int) *Token {
	return &Token{type_: type_, literal: literal, line: line, column: column}
}

func (this *Token) String() string {
	return fmt.Sprintf("Token[type: %d, literal: \"%s\", line: %d, column: %d]", this.type_, this.literal, this.line, this.column)
}

func (this *Token) ErrorMsg() string {
	return fmt.Sprintf("at \"%s\", line: %d, column: %d]", this.literal, this.line, this.column)
}

type TokenType int

const (
	GoLexerNone                   TokenType = 0
	GoLexerBREAK                  TokenType = 1
	GoLexerDEFAULT                TokenType = 2
	GoLexerFUNC                   TokenType = 3
	GoLexerINTERFACE              TokenType = 4
	GoLexerSELECT                 TokenType = 5
	GoLexerCASE                   TokenType = 6
	GoLexerDEFER                  TokenType = 7
	GoLexerGO                     TokenType = 8
	GoLexerMAP                    TokenType = 9
	GoLexerSTRUCT                 TokenType = 10
	GoLexerCHAN                   TokenType = 11
	GoLexerELSE                   TokenType = 12
	GoLexerGOTO                   TokenType = 13
	GoLexerPACKAGE                TokenType = 14
	GoLexerSWITCH                 TokenType = 15
	GoLexerCONST                  TokenType = 16
	GoLexerFALLTHROUGH            TokenType = 17
	GoLexerIF                     TokenType = 18
	GoLexerRANGE                  TokenType = 19
	GoLexerTYPE                   TokenType = 20
	GoLexerCONTINUE               TokenType = 21
	GoLexerFOR                    TokenType = 22
	GoLexerIMPORT                 TokenType = 23
	GoLexerRETURN                 TokenType = 24
	GoLexerVAR                    TokenType = 25
	GoLexerNIL_LIT                TokenType = 26
	GoLexerIDENTIFIER             TokenType = 27
	GoLexerL_PAREN                TokenType = 28
	GoLexerR_PAREN                TokenType = 29
	GoLexerL_CURLY                TokenType = 30
	GoLexerR_CURLY                TokenType = 31
	GoLexerL_BRACKET              TokenType = 32
	GoLexerR_BRACKET              TokenType = 33
	GoLexerASSIGN                 TokenType = 34
	GoLexerCOMMA                  TokenType = 35
	GoLexerSEMI                   TokenType = 36
	GoLexerCOLON                  TokenType = 37
	GoLexerDOT                    TokenType = 38
	GoLexerPLUS_PLUS              TokenType = 39
	GoLexerMINUS_MINUS            TokenType = 40
	GoLexerDECLARE_ASSIGN         TokenType = 41
	GoLexerELLIPSIS               TokenType = 42
	GoLexerLOGICAL_OR             TokenType = 43
	GoLexerLOGICAL_AND            TokenType = 44
	GoLexerEQUALS                 TokenType = 45
	GoLexerNOT_EQUALS             TokenType = 46
	GoLexerLESS                   TokenType = 47
	GoLexerLESS_OR_EQUALS         TokenType = 48
	GoLexerGREATER                TokenType = 49
	GoLexerGREATER_OR_EQUALS      TokenType = 50
	GoLexerOR                     TokenType = 51
	GoLexerDIV                    TokenType = 52
	GoLexerMOD                    TokenType = 53
	GoLexerLSHIFT                 TokenType = 54
	GoLexerRSHIFT                 TokenType = 55
	GoLexerBIT_CLEAR              TokenType = 56
	GoLexerEXCLAMATION            TokenType = 57
	GoLexerPLUS                   TokenType = 58
	GoLexerMINUS                  TokenType = 59
	GoLexerCARET                  TokenType = 60
	GoLexerSTAR                   TokenType = 61
	GoLexerAMPERSAND              TokenType = 62
	GoLexerRECEIVE                TokenType = 63
	GoLexerDECIMAL_LIT            TokenType = 64
	GoLexerBINARY_LIT             TokenType = 65
	GoLexerOCTAL_LIT              TokenType = 66
	GoLexerHEX_LIT                TokenType = 67
	GoLexerFLOAT_LIT              TokenType = 68
	GoLexerDECIMAL_FLOAT_LIT      TokenType = 69
	GoLexerHEX_FLOAT_LIT          TokenType = 70
	GoLexerIMAGINARY_LIT          TokenType = 71
	GoLexerRUNE_LIT               TokenType = 72
	GoLexerBYTE_VALUE             TokenType = 73
	GoLexerOCTAL_BYTE_VALUE       TokenType = 74
	GoLexerHEX_BYTE_VALUE         TokenType = 75
	GoLexerLITTLE_U_VALUE         TokenType = 76
	GoLexerBIG_U_VALUE            TokenType = 77
	GoLexerRAW_STRING_LIT         TokenType = 78
	GoLexerINTERPRETED_STRING_LIT TokenType = 79
	GoLexerWS                     TokenType = 80
	GoLexerANNOTATION_COMMENT     TokenType = 81
	GoLexerCOMMENT                TokenType = 82
	GoLexerTERMINATOR             TokenType = 83
	GoLexerLINE_COMMENT           TokenType = 84
	GoLexerWS_NLSEMI              TokenType = 85
	GoLexerCOMMENT_NLSEMI         TokenType = 86
	GoLexerLINE_COMMENT_NLSEMI    TokenType = 87
	GoLexerEOS                    TokenType = 88
	GoLexerOTHER                  TokenType = 89
)
