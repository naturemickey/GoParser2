package lex

import "unicode"

type nfa struct {
	start  *state
	finish *state
}

var NFA = GoLexerNFA()

func (this *nfa) SetType(type_ TokenType) {
	this.finish.SetType(type_)
}

func And(nfas ...*nfa) *nfa {
	var start *nfa
	var n *nfa
	for i, nfa := range nfas {
		if i == 0 {
			start = nfa
		} else {
			n.finish.addCharPath(e, nfa.start)
		}
		n = nfa
	}
	return &nfa{start: start.start, finish: n.finish}
}

func Or(nfas ...*nfa) *nfa {
	start := NewState()
	finish := NewState()

	for _, nfa := range nfas {
		start.addCharPath(e, nfa.start)
		nfa.finish.addCharPath(e, finish)
	}

	return &nfa{start, finish}
}

// 0个或更多
func Kc(nfa *nfa) *nfa {
	nfa.start.addCharPath(e, nfa.finish)
	nfa.finish.addCharPath(e, nfa.start)
	return nfa
}

// 1个或更多
func Kc1(nfa *nfa) *nfa {
	nfa.finish.addCharPath(e, nfa.start)
	return nfa
}

// 0个或1个
func Le1(nfa *nfa) *nfa {
	nfa.start.addCharPath(e, nfa.finish)
	return nfa
}

func NewNfaWithChar(char rune, type_ TokenType) *nfa {
	start := NewState()
	finish := NewStateWithType(type_)
	start.addCharPath(char, finish)
	return &nfa{start, finish}
}

// str每个字符and
func NewNfaWithString(str string, type_ TokenType) *nfa {
	var nfas []*nfa
	var chars = []rune(str)
	for _, char := range chars {
		nfas = append(nfas, NewNfaWithChar(char, GoLexerNone))
	}
	nfa := And(nfas...)
	nfa.SetType(type_)
	return nfa
}

// str每个字符or
func NewNfaWithChars(str string, type_ TokenType) *nfa {
	var nfas []*nfa
	for _, char := range []rune(str) {
		nfas = append(nfas, NewNfaWithChar(char, type_))
	}
	return Or(nfas...)
}

func NewNfaWithRegular(regular func(rune) bool, type_ TokenType) *nfa {
	start := NewState()
	finish := NewStateWithType(type_)
	start.addRegularPath(regular, finish)
	return &nfa{start, finish}
}

func GoLexerNFA() *nfa {
	var nfas = []*nfa{
		// Keywords
		NewNfaWithString("break", GoLexerBREAK),
		NewNfaWithString("default", GoLexerDEFAULT),
		NewNfaWithString("func", GoLexerFUNC),
		NewNfaWithString("interface", GoLexerINTERFACE),
		NewNfaWithString("select", GoLexerSELECT),
		NewNfaWithString("case", GoLexerCASE),
		NewNfaWithString("defer", GoLexerDEFER),
		NewNfaWithString("go", GoLexerGO),
		NewNfaWithString("map", GoLexerMAP),
		NewNfaWithString("struct", GoLexerSTRUCT),
		NewNfaWithString("chan", GoLexerCHAN),
		NewNfaWithString("else", GoLexerELSE),
		NewNfaWithString("goto", GoLexerGOTO),
		NewNfaWithString("package", GoLexerPACKAGE),
		NewNfaWithString("switch", GoLexerSWITCH),
		NewNfaWithString("const", GoLexerCONST),
		NewNfaWithString("fallthrough", GoLexerFALLTHROUGH),
		NewNfaWithString("if", GoLexerIF),
		NewNfaWithString("range", GoLexerRANGE),
		NewNfaWithString("type", GoLexerTYPE),
		NewNfaWithString("continue", GoLexerCONTINUE),
		NewNfaWithString("for", GoLexerFOR),
		NewNfaWithString("import", GoLexerIMPORT),
		NewNfaWithString("return", GoLexerRETURN),
		NewNfaWithString("var", GoLexerVAR),
		NewNfaWithString("nil", GoLexerNIL_LIT),
		// IDENTIFIER
		NewIdentifierNfa(),
		// Punctuation
		NewNfaWithString("(", GoLexerL_PAREN),
		NewNfaWithString(")", GoLexerR_PAREN),
		NewNfaWithString("{", GoLexerL_CURLY),
		NewNfaWithString("}", GoLexerR_CURLY),
		NewNfaWithString("[", GoLexerL_BRACKET),
		NewNfaWithString("]", GoLexerR_BRACKET),
		NewNfaWithString("=", GoLexerASSIGN),
		NewNfaWithString(",", GoLexerCOMMA),
		NewNfaWithString(";", GoLexerSEMI),
		NewNfaWithString(":", GoLexerCOLON),
		NewNfaWithString(".", GoLexerDOT),
		NewNfaWithString("++", GoLexerPLUS_PLUS),
		NewNfaWithString("--", GoLexerMINUS_MINUS),
		NewNfaWithString(":=", GoLexerDECLARE_ASSIGN),
		NewNfaWithString("...", GoLexerELLIPSIS),
		// Logical
		NewNfaWithString("||", GoLexerLOGICAL_OR),
		NewNfaWithString("&&", GoLexerLOGICAL_AND),
		// Relation operators
		NewNfaWithString("==", GoLexerEQUALS),
		NewNfaWithString("!=", GoLexerNOT_EQUALS),
		NewNfaWithString("<", GoLexerLESS),
		NewNfaWithString("<=", GoLexerLESS_OR_EQUALS),
		NewNfaWithString(">", GoLexerGREATER),
		NewNfaWithString(">=", GoLexerGREATER_OR_EQUALS),
		// Arithmetic operators
		NewNfaWithString("|", GoLexerOR),
		NewNfaWithString("/", GoLexerDIV),
		NewNfaWithString("%", GoLexerMOD),
		NewNfaWithString("<<", GoLexerLSHIFT),
		NewNfaWithString(">>", GoLexerRSHIFT),
		NewNfaWithString("&^", GoLexerBIT_CLEAR),
		// Unary operators
		NewNfaWithString("!", GoLexerEXCLAMATION),
		// Mixed operators
		NewNfaWithString("+", GoLexerPLUS),
		NewNfaWithString("-", GoLexerMINUS),
		NewNfaWithString("^", GoLexerCARET),
		NewNfaWithString("*", GoLexerSTAR),
		NewNfaWithString("&", GoLexerAMPERSAND),
		NewNfaWithString("<-", GoLexerRECEIVE),
		// Number literals
		New_DECIMAL_LIT_nfa(),
		New_BINARY_LIT_nfa(),
		New_OCTAL_LIT_nfa(),
		New_HEX_LIT_nfa(),
		New_FLOAT_LIT_nfa(),
		New_DECIMAL_FLOAT_LIT_nfa(),
		New_HEX_FLOAT_LIT_nfa(),
		New_IMAGINARY_LIT_nfa(),
		New_RUNE_LIT_nfa(),
		New_BYTE_VALUE_nfa(),
		New_OCTAL_BYTE_VALUE_nfa(),
		New_HEX_BYTE_VALUE_nfa(),
		New_LITTLE_U_VALUE_nfa(),
		New_BIG_U_VALUE_nfa(),
		// String literals
		New_RAW_STRING_LIT_nfa(),
		New_INTERPRETED_STRING_LIT_nfa(),
		// Hidden tokens
		New_WS_nfa(),
		New_ANNOTATION_COMMENT_nfa(), // 这个不是被忽略的，放在这里好看一点
		New_COMMENT_nfa(),
		New_TERMINATOR_nfa(),
		New_LINE_COMMENT_nfa(),
	}
	return Or(nfas...)
}

func New_WS_nfa() *nfa {
	// [ \t]+
	return Kc1(NewNfaWithChars(" \t", GoLexerWS))
}

func New_TERMINATOR_nfa() *nfa {
	// [\r\n]+
	return Kc1(NewNfaWithChars("\r\n", GoLexerTERMINATOR))
}

func New_ANNOTATION_COMMENT_nfa() *nfa {
	// '/*@\w+' ('(' \w+=\w+ (',' \w+=\w+)* ')')? '*/' ;
	nfa := And(
		NewNfaWithString("/*@", GoLexerNone),
		new_w_plus_nfa(),
		Le1(And(
			NewNfaWithString("(", GoLexerNone),
			new_w_plus_nfa(), NewNfaWithString("=", GoLexerNone), new_w_plus_nfa(),
			Kc(And(
				NewNfaWithString(",", GoLexerNone),
				new_w_plus_nfa(), NewNfaWithString("=", GoLexerNone), new_w_plus_nfa(),
			)),
			NewNfaWithString(")", GoLexerNone),
		)),
		NewNfaWithString("*/", GoLexerNone),
	)
	nfa.SetType(GoLexerANNOTATION_COMMENT)
	return nfa
}

func new_w_plus_nfa() *nfa {
	// \w+
	return Kc1(NewNfaWithRegular(func(c rune) bool {
		return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '_'
	}, GoLexerNone))
}

func New_COMMENT_nfa() *nfa {
	// '/*' .*? '*/'
	// todo 这个地方要看一下如何防止贪婪匹配
	commitStart := NewNfaWithString("/*", GoLexerNone)

	stateToBeFinish := NewState()
	stateAny := NewState()
	stateFinish := NewStateWithType(GoLexerCOMMENT)

	commitStart.finish.addCharPath('*', stateToBeFinish)
	commitStart.finish.addRegularPath(func(c rune) bool {
		return c != '*'
	}, stateAny)
	stateAny.addCharPath('*', stateToBeFinish)
	stateAny.addRegularPath(func(c rune) bool {
		return c != '*'
	}, stateAny)
	stateToBeFinish.addRegularPath(func(c rune) bool {
		return c != '/'
	}, stateAny)
	stateToBeFinish.addCharPath('/', stateFinish)

	//nfa := And(
	//	commitStart,
	//	Le1(Kc(
	//		NewNfaWithRegular(func(c rune) bool { return true }, GoLexerNone),
	//	)),
	//	NewNfaWithString("*/", GoLexerNone),
	//)
	//nfa.SetType(GoLexerCOMMENT)
	return &nfa{commitStart.start, stateFinish}
}

func New_LINE_COMMENT_nfa() *nfa {
	// '//' ~[\r\n]*
	nfa := And(
		NewNfaWithString("//", GoLexerNone),
		Kc(NewNfaWithRegular(func(c rune) bool {
			return c != '\r' && c != '\n'
		}, GoLexerNone)),
	)
	nfa.SetType(GoLexerLINE_COMMENT)
	return nfa
}

func New_RAW_STRING_LIT_nfa() *nfa {
	// '`' ~'`'* '`'
	nfa := And(
		NewNfaWithChar('`', GoLexerNone),
		Kc(NewNfaWithRegular(func(c rune) bool {
			return c != '`'
		}, GoLexerNone)),
		NewNfaWithChar('`', GoLexerNone),
	)
	nfa.SetType(GoLexerRAW_STRING_LIT)
	return nfa
}

func New_INTERPRETED_STRING_LIT_nfa() *nfa {
	// '"' (~["\\] | ESCAPED_VALUE)* '"'
	nfa := And(
		NewNfaWithChar('"', GoLexerNone),
		Kc(Or(
			NewNfaWithRegular(func(c rune) bool {
				return c != '"' && c != '\\'
			}, GoLexerNone),
			new_ESCAPED_VALUE_nfa(),
		)),
		NewNfaWithChar('"', GoLexerNone),
	)
	nfa.SetType(GoLexerINTERPRETED_STRING_LIT)
	return nfa
}

func NewIdentifierNfa() *nfa {
	// LETTER (LETTER | UNICODE_DIGIT)*
	nfa := And(
		new_LETTER_nfa(),
		Kc(Or(
			new_LETTER_nfa(),
			new_UNICODE_DIGIT_nfa(),
		)),
	)
	nfa.SetType(GoLexerIDENTIFIER)
	return nfa
}
func New_DECIMAL_LIT_nfa() *nfa {
	// ('0' | [1-9] ('_'? [0-9])*)

	// '0' :
	nfa0 := NewNfaWithChar('0', GoLexerNone)
	// [1-9] :
	nfa1 := NewNfaWithRegular(func(char rune) bool {
		return '1' <= char && char <= '9'
	}, GoLexerNone)

	// '_'? :
	nfa2 := NewNfaWithChar('_', GoLexerNone)
	nfa2 = Le1(nfa2)

	// [0-9] :
	nfa3 := NewNfaWithRegular(func(char rune) bool {
		return '0' <= char && char <= '9'
	}, GoLexerNone)

	// ('_'? [0-9])*
	nfa4 := And(nfa2, nfa3)
	nfa4 = Kc(nfa4)

	nfa := Or(nfa0, And(nfa1, nfa4))
	nfa.SetType(GoLexerDECIMAL_LIT)
	return nfa
}
func New_BINARY_LIT_nfa() *nfa {
	// '0' [bB] ('_'? [01])+

	// '_'? :
	nfa_ := Le1(NewNfaWithChar('_', GoLexerNone))

	// ('_'? [01])+
	nfa := Kc1(And(nfa_, NewNfaWithChars("01", GoLexerNone)))

	nfa = And(
		NewNfaWithChar('0', GoLexerNone),
		NewNfaWithChars("bB", GoLexerNone),
		nfa,
	)
	nfa.SetType(GoLexerBINARY_LIT)
	return nfa
}
func New_OCTAL_LIT_nfa() *nfa {
	// '0' [oO]? ('_'? [0-7])+
	nfa := Or(
		NewNfaWithChar('0', GoLexerNone),
		Le1(NewNfaWithChars("oO", GoLexerNone)),
		Kc1(And(
			Le1(NewNfaWithChar('_', GoLexerNone)),
			NewNfaWithRegular(func(c rune) bool {
				return '0' <= c && c <= '7'
			}, GoLexerNone),
		)),
	)
	nfa.SetType(GoLexerOCTAL_LIT)
	return nfa
}
func New_HEX_LIT_nfa() *nfa {
	// '0' [xX]  ('_'? HEX_DIGIT)+
	nfa := And(
		NewNfaWithChar('0', GoLexerNone),
		NewNfaWithChars("xX", GoLexerNone),
		Kc1(And(
			Le1(NewNfaWithChar('_', GoLexerNone)),
			new_HEX_DIGIT_nfa(),
		)),
	)
	nfa.SetType(GoLexerHEX_LIT)
	return nfa
}
func new_HEX_DIGIT_nfa() *nfa {
	// [0-9a-fA-F]
	return NewNfaWithRegular(func(c rune) bool {
		return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
	}, GoLexerNone)
}
func New_FLOAT_LIT_nfa() *nfa {
	// DECIMAL_FLOAT_LIT | HEX_FLOAT_LIT
	nfa := Or(
		New_DECIMAL_FLOAT_LIT_nfa(),
		New_HEX_FLOAT_LIT_nfa(),
	)
	nfa.SetType(GoLexerFLOAT_LIT)
	return nfa
}
func New_DECIMAL_FLOAT_LIT_nfa() *nfa {
	// DECIMALS ('.' DECIMALS? EXPONENT? | EXPONENT)
	// | '.' DECIMALS EXPONENT?
	nfa := Or(
		And(
			new_DECIMALS_nfa(),
			Or(
				And(
					NewNfaWithChar('.', GoLexerNone),
					Le1(new_DECIMALS_nfa()),
					Le1(new_EXPONENT_nfa()),
				),
				new_EXPONENT_nfa(),
			),
		),
		And(
			NewNfaWithChar('.', GoLexerNone),
			new_DECIMALS_nfa(),
			Le1(new_EXPONENT_nfa()),
		),
	)
	nfa.SetType(GoLexerDECIMAL_FLOAT_LIT)
	return nfa
}
func New_HEX_FLOAT_LIT_nfa() *nfa {
	// HEX_FLOAT_LIT : '0' [xX] HEX_MANTISSA HEX_EXPONENT
	// HEX_MANTISSA  : ('_'? HEX_DIGIT)+ ('.' ('_'? HEX_DIGIT)*)?
	//               | '.' HEX_DIGIT ('_'? HEX_DIGIT)*;
	HEX_MANTISSA := Or(
		// ('_'? HEX_DIGIT)+ ('.' ('_'? HEX_DIGIT)*)?
		And(
			// ('_'? HEX_DIGIT)+
			Kc1(And(
				Le1(NewNfaWithChar('_', GoLexerNone)),
				new_HEX_DIGIT_nfa(),
			)),
			// ('.' ( '_'? HEX_DIGIT)*)?
			Le1(And(
				NewNfaWithChar('.', GoLexerNone),
				Kc(And(
					Le1(NewNfaWithChar('_', GoLexerNone)),
					new_HEX_DIGIT_nfa(),
				)),
			)),
		),
		// '.' HEX_DIGIT ('_'? HEX_DIGIT)*
		And(
			NewNfaWithChar('.', GoLexerNone),
			new_HEX_DIGIT_nfa(),
			Kc(And(
				Le1(NewNfaWithChar('_', GoLexerNone)),
				new_HEX_DIGIT_nfa(),
			)),
		),
	)

	// HEX_EXPONENT  : [pP] [+-]? DECIMALS;
	HEX_EXPONENT := And(
		NewNfaWithChars("pP", GoLexerNone),
		Le1(NewNfaWithChars("+-", GoLexerNone)),
		new_DECIMALS_nfa(),
	)

	// '0' [xX] HEX_MANTISSA HEX_EXPONENT
	nfa := And(
		NewNfaWithChar('0', GoLexerNone),
		NewNfaWithChars("xX", GoLexerNone),
		HEX_MANTISSA,
		HEX_EXPONENT,
	)
	nfa.SetType(GoLexerHEX_FLOAT_LIT)
	return nfa
}
func New_IMAGINARY_LIT_nfa() *nfa {
	// (DECIMAL_LIT | BINARY_LIT | OCTAL_LIT | HEX_LIT | FLOAT_LIT) 'i'
	nfa := And(
		Or(
			New_DECIMAL_LIT_nfa(),
			New_BINARY_LIT_nfa(),
			New_OCTAL_LIT_nfa(),
			New_HEX_LIT_nfa(),
			New_FLOAT_LIT_nfa(),
		),
		NewNfaWithChar('i', GoLexerNone),
	)
	nfa.SetType(GoLexerIMAGINARY_LIT)
	return nfa
}
func New_RUNE_LIT_nfa() *nfa {
	// '\'' (UNICODE_VALUE | BYTE_VALUE) '\''
	nfa := And(
		NewNfaWithChar('\'', GoLexerNone),
		Or(
			new_UNICODE_VALUE_nfa(),
			New_BYTE_VALUE_nfa(),
		),
		NewNfaWithChar('\'', GoLexerNone),
	)
	nfa.SetType(GoLexerRUNE_LIT)
	return nfa
}
func New_BYTE_VALUE_nfa() *nfa {
	// OCTAL_BYTE_VALUE | HEX_BYTE_VALUE
	nfa := Or(
		New_OCTAL_BYTE_VALUE_nfa(),
		New_HEX_BYTE_VALUE_nfa(),
	)
	nfa.SetType(GoLexerBYTE_VALUE)
	return nfa
}
func New_OCTAL_BYTE_VALUE_nfa() *nfa {
	// '\\' OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
	nfa := And(
		NewNfaWithChar('\\', GoLexerNone),
		new_OCTAL_DIGIT_nfa(),
		new_OCTAL_DIGIT_nfa(),
		new_OCTAL_DIGIT_nfa(),
	)
	nfa.SetType(GoLexerOCTAL_BYTE_VALUE)
	return nfa
}
func New_HEX_BYTE_VALUE_nfa() *nfa {
	// '\\' 'x'  HEX_DIGIT HEX_DIGIT
	nfa := And(
		NewNfaWithChar('\\', GoLexerNone),
		NewNfaWithChar('x', GoLexerNone),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
	)
	nfa.SetType(GoLexerHEX_BYTE_VALUE)
	return nfa
}
func New_LITTLE_U_VALUE_nfa() *nfa {
	// '\\' 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
	nfa := And(
		NewNfaWithChar('\\', GoLexerNone),
		NewNfaWithChar('u', GoLexerNone),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
	)
	nfa.SetType(GoLexerLITTLE_U_VALUE)
	return nfa
}
func New_BIG_U_VALUE_nfa() *nfa {
	// '\\' 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
	nfa := And(
		NewNfaWithChar('\\', GoLexerNone),
		NewNfaWithChar('U', GoLexerNone),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
		new_HEX_DIGIT_nfa(),
	)
	nfa.SetType(GoLexerBIG_U_VALUE)
	return nfa
}

func new_DECIMALS_nfa() *nfa {
	// [0-9] ('_'? [0-9])*
	return And(
		NewNfaWithRegular(func(c rune) bool {
			return '0' <= c && c <= '9'
		}, GoLexerNone),
		Kc(And(
			Le1(NewNfaWithChar('_', GoLexerNone)),
			NewNfaWithRegular(func(c rune) bool {
				return '0' <= c && c <= '9'
			}, GoLexerNone),
		)),
	)
}

func new_EXPONENT_nfa() *nfa {
	// [eE] [+-]? DECIMALS
	return And(
		NewNfaWithChars("eE", GoLexerNone),
		Le1(NewNfaWithChars("+-", GoLexerNone)),
		new_DECIMALS_nfa(),
	)
}

func new_OCTAL_DIGIT_nfa() *nfa {
	return NewNfaWithRegular(func(c rune) bool {
		return '0' <= c && c <= '7'
	}, GoLexerNone)
}

func new_UNICODE_VALUE_nfa() *nfa {
	// ~[\r\n'] | LITTLE_U_VALUE | BIG_U_VALUE | ESCAPED_VALUE
	return Or(
		NewNfaWithRegular(func(c rune) bool {
			return c != '\r' && c != '\n' && c != '\''
		}, GoLexerNone),
		New_LITTLE_U_VALUE_nfa(),
		New_BIG_U_VALUE_nfa(),
		new_ESCAPED_VALUE_nfa(),
	)
}

func new_ESCAPED_VALUE_nfa() *nfa {
	// '\\' ('u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
	// | 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
	// | [abfnrtv\\'"]
	// | OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
	// | 'x' HEX_DIGIT HEX_DIGIT)
	return And(
		NewNfaWithChar('\\', GoLexerNone),
		Or(
			// 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
			And(
				NewNfaWithChar('u', GoLexerNone),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
			),
			// 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
			And(
				NewNfaWithChar('U', GoLexerNone),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
			),
			// [abfnrtv\\'"]
			NewNfaWithChars("abfnrtv\\\\'\"", GoLexerNone),
			// OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
			And(
				new_OCTAL_DIGIT_nfa(),
				new_OCTAL_DIGIT_nfa(),
				new_OCTAL_DIGIT_nfa(),
			),
			// 'x' HEX_DIGIT HEX_DIGIT
			And(
				NewNfaWithChar('x', GoLexerNone),
				new_HEX_DIGIT_nfa(),
				new_HEX_DIGIT_nfa(),
			),
		),
	)
}

func new_UNICODE_DIGIT_nfa() *nfa {
	return NewNfaWithRegular(func(c rune) bool {
		return unicode.IsDigit(c)
	}, GoLexerNone)
}

func new_LETTER_nfa() *nfa {
	// UNICODE_LETTER | '_'
	return Or(new_UNICODE_LETTER_nfa(), NewNfaWithChar('_', GoLexerNone))
}

func new_UNICODE_LETTER_nfa() *nfa {
	return NewNfaWithRegular(func(c rune) bool {
		return unicode.IsLetter(c)
	}, GoLexerNone)
}
