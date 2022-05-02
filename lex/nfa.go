package lex

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

func New_COMMENT_nfa() *nfa {
	// '/*' .*? '*/'
	nfa := And(
		NewNfaWithChars("/*", GoLexerNone),
		Le1(Kc(
			NewNfaWithRegular(func(c rune) bool { return true }, GoLexerNone),
		)),
		NewNfaWithChars("*/", GoLexerNone),
	)
	nfa.SetType(GoLexerCOMMENT)
	return nfa
}

func New_LINE_COMMENT_nfa() *nfa {
	// '//' ~[\r\n]*
	nfa := And(
		NewNfaWithChars("//", GoLexerNone),
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
			NewNfaWithString("abfnrtv\\\\'\"", GoLexerNone),
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
	/* [\u0030-\u0039]
	| [\u0660-\u0669]
	| [\u06F0-\u06F9]
	| [\u0966-\u096F]
	| [\u09E6-\u09EF]
	| [\u0A66-\u0A6F]
	| [\u0AE6-\u0AEF]
	| [\u0B66-\u0B6F]
	| [\u0BE7-\u0BEF]
	| [\u0C66-\u0C6F]
	| [\u0CE6-\u0CEF]
	| [\u0D66-\u0D6F]
	| [\u0E50-\u0E59]
	| [\u0ED0-\u0ED9]
	| [\u0F20-\u0F29]
	| [\u1040-\u1049]
	| [\u1369-\u1371]
	| [\u17E0-\u17E9]
	| [\u1810-\u1819]
	| [\uFF10-\uFF19]*/
	return NewNfaWithRegular(func(c rune) bool {
		return ('\u0030' <= c && c <= '\u0039') ||
			('\u0660' <= c && c <= '\u0669') ||
			('\u06F0' <= c && c <= '\u06F9') ||
			('\u0966' <= c && c <= '\u096F') ||
			('\u09E6' <= c && c <= '\u09EF') ||
			('\u0A66' <= c && c <= '\u0A6F') ||
			('\u0AE6' <= c && c <= '\u0AEF') ||
			('\u0B66' <= c && c <= '\u0B6F') ||
			('\u0BE7' <= c && c <= '\u0BEF') ||
			('\u0C66' <= c && c <= '\u0C6F') ||
			('\u0CE6' <= c && c <= '\u0CEF') ||
			('\u0D66' <= c && c <= '\u0D6F') ||
			('\u0E50' <= c && c <= '\u0E59') ||
			('\u0ED0' <= c && c <= '\u0ED9') ||
			('\u0F20' <= c && c <= '\u0F29') ||
			('\u1040' <= c && c <= '\u1049') ||
			('\u1369' <= c && c <= '\u1371') ||
			('\u17E0' <= c && c <= '\u17E9') ||
			('\u1810' <= c && c <= '\u1819') ||
			('\uFF10' <= c && c <= '\uFF19')
	}, GoLexerNone)
}

func new_LETTER_nfa() *nfa {
	// UNICODE_LETTER | '_'
	return Or(new_UNICODE_LETTER_nfa(), NewNfaWithChar('_', GoLexerNone))
}

func new_UNICODE_LETTER_nfa() *nfa {
	/* [\u0041-\u005A]
	   | [\u0061-\u007A]
	   | [\u00AA]
	   | [\u00B5]
	   | [\u00BA]
	   | [\u00C0-\u00D6]
	   | [\u00D8-\u00F6]
	   | [\u00F8-\u021F]
	   | [\u0222-\u0233]
	   | [\u0250-\u02AD]
	   | [\u02B0-\u02B8]
	   | [\u02BB-\u02C1]
	   | [\u02D0-\u02D1]
	   | [\u02E0-\u02E4]
	   | [\u02EE]
	   | [\u037A]
	   | [\u0386]
	   | [\u0388-\u038A]
	   | [\u038C]
	   | [\u038E-\u03A1]
	   | [\u03A3-\u03CE]
	   | [\u03D0-\u03D7]
	   | [\u03DA-\u03F3]
	   | [\u0400-\u0481]
	   | [\u048C-\u04C4]
	   | [\u04C7-\u04C8]
	   | [\u04CB-\u04CC]
	   | [\u04D0-\u04F5]
	   | [\u04F8-\u04F9]
	   | [\u0531-\u0556]
	   | [\u0559]
	   | [\u0561-\u0587]
	   | [\u05D0-\u05EA]
	   | [\u05F0-\u05F2]
	   | [\u0621-\u063A]
	   | [\u0640-\u064A]
	   | [\u0671-\u06D3]
	   | [\u06D5]
	   | [\u06E5-\u06E6]
	   | [\u06FA-\u06FC]
	   | [\u0710]
	   | [\u0712-\u072C]
	   | [\u0780-\u07A5]
	   | [\u0905-\u0939]
	   | [\u093D]
	   | [\u0950]
	   | [\u0958-\u0961]
	   | [\u0985-\u098C]
	   | [\u098F-\u0990]
	   | [\u0993-\u09A8]
	   | [\u09AA-\u09B0]
	   | [\u09B2]
	   | [\u09B6-\u09B9]
	   | [\u09DC-\u09DD]
	   | [\u09DF-\u09E1]
	   | [\u09F0-\u09F1]
	   | [\u0A05-\u0A0A]
	   | [\u0A0F-\u0A10]
	   | [\u0A13-\u0A28]
	   | [\u0A2A-\u0A30]
	   | [\u0A32-\u0A33]
	   | [\u0A35-\u0A36]
	   | [\u0A38-\u0A39]
	   | [\u0A59-\u0A5C]
	   | [\u0A5E]
	   | [\u0A72-\u0A74]
	   | [\u0A85-\u0A8B]
	   | [\u0A8D]
	   | [\u0A8F-\u0A91]
	   | [\u0A93-\u0AA8]
	   | [\u0AAA-\u0AB0]
	   | [\u0AB2-\u0AB3]
	   | [\u0AB5-\u0AB9]
	   | [\u0ABD]
	   | [\u0AD0]
	   | [\u0AE0]
	   | [\u0B05-\u0B0C]
	   | [\u0B0F-\u0B10]
	   | [\u0B13-\u0B28]
	   | [\u0B2A-\u0B30]
	   | [\u0B32-\u0B33]
	   | [\u0B36-\u0B39]
	   | [\u0B3D]
	   | [\u0B5C-\u0B5D]
	   | [\u0B5F-\u0B61]
	   | [\u0B85-\u0B8A]
	   | [\u0B8E-\u0B90]
	   | [\u0B92-\u0B95]
	   | [\u0B99-\u0B9A]
	   | [\u0B9C]
	   | [\u0B9E-\u0B9F]
	   | [\u0BA3-\u0BA4]
	   | [\u0BA8-\u0BAA]
	   | [\u0BAE-\u0BB5]
	   | [\u0BB7-\u0BB9]
	   | [\u0C05-\u0C0C]
	   | [\u0C0E-\u0C10]
	   | [\u0C12-\u0C28]
	   | [\u0C2A-\u0C33]
	   | [\u0C35-\u0C39]
	   | [\u0C60-\u0C61]
	   | [\u0C85-\u0C8C]
	   | [\u0C8E-\u0C90]
	   | [\u0C92-\u0CA8]
	   | [\u0CAA-\u0CB3]
	   | [\u0CB5-\u0CB9]
	   | [\u0CDE]
	   | [\u0CE0-\u0CE1]
	   | [\u0D05-\u0D0C]
	   | [\u0D0E-\u0D10]
	   | [\u0D12-\u0D28]
	   | [\u0D2A-\u0D39]
	   | [\u0D60-\u0D61]
	   | [\u0D85-\u0D96]
	   | [\u0D9A-\u0DB1]
	   | [\u0DB3-\u0DBB]
	   | [\u0DBD]
	   | [\u0DC0-\u0DC6]
	   | [\u0E01-\u0E30]
	   | [\u0E32-\u0E33]
	   | [\u0E40-\u0E46]
	   | [\u0E81-\u0E82]
	   | [\u0E84]
	   | [\u0E87-\u0E88]
	   | [\u0E8A]
	   | [\u0E8D]
	   | [\u0E94-\u0E97]
	   | [\u0E99-\u0E9F]
	   | [\u0EA1-\u0EA3]
	   | [\u0EA5]
	   | [\u0EA7]
	   | [\u0EAA-\u0EAB]
	   | [\u0EAD-\u0EB0]
	   | [\u0EB2-\u0EB3]
	   | [\u0EBD-\u0EC4]
	   | [\u0EC6]
	   | [\u0EDC-\u0EDD]
	   | [\u0F00]
	   | [\u0F40-\u0F6A]
	   | [\u0F88-\u0F8B]
	   | [\u1000-\u1021]
	   | [\u1023-\u1027]
	   | [\u1029-\u102A]
	   | [\u1050-\u1055]
	   | [\u10A0-\u10C5]
	   | [\u10D0-\u10F6]
	   | [\u1100-\u1159]
	   | [\u115F-\u11A2]
	   | [\u11A8-\u11F9]
	   | [\u1200-\u1206]
	   | [\u1208-\u1246]
	   | [\u1248]
	   | [\u124A-\u124D]
	   | [\u1250-\u1256]
	   | [\u1258]
	   | [\u125A-\u125D]
	   | [\u1260-\u1286]
	   | [\u1288]
	   | [\u128A-\u128D]
	   | [\u1290-\u12AE]
	   | [\u12B0]
	   | [\u12B2-\u12B5]
	   | [\u12B8-\u12BE]
	   | [\u12C0]
	   | [\u12C2-\u12C5]
	   | [\u12C8-\u12CE]
	   | [\u12D0-\u12D6]
	   | [\u12D8-\u12EE]
	   | [\u12F0-\u130E]
	   | [\u1310]
	   | [\u1312-\u1315]
	   | [\u1318-\u131E]
	   | [\u1320-\u1346]
	   | [\u1348-\u135A]
	   | [\u13A0-\u13B0]
	   | [\u13B1-\u13F4]
	   | [\u1401-\u1676]
	   | [\u1681-\u169A]
	   | [\u16A0-\u16EA]
	   | [\u1780-\u17B3]
	   | [\u1820-\u1877]
	   | [\u1880-\u18A8]
	   | [\u1E00-\u1E9B]
	   | [\u1EA0-\u1EE0]
	   | [\u1EE1-\u1EF9]
	   | [\u1F00-\u1F15]
	   | [\u1F18-\u1F1D]
	   | [\u1F20-\u1F39]
	   | [\u1F3A-\u1F45]
	   | [\u1F48-\u1F4D]
	   | [\u1F50-\u1F57]
	   | [\u1F59]
	   | [\u1F5B]
	   | [\u1F5D]
	   | [\u1F5F-\u1F7D]
	   | [\u1F80-\u1FB4]
	   | [\u1FB6-\u1FBC]
	   | [\u1FBE]
	   | [\u1FC2-\u1FC4]
	   | [\u1FC6-\u1FCC]
	   | [\u1FD0-\u1FD3]
	   | [\u1FD6-\u1FDB]
	   | [\u1FE0-\u1FEC]
	   | [\u1FF2-\u1FF4]
	   | [\u1FF6-\u1FFC]
	   | [\u207F]
	   | [\u2102]
	   | [\u2107]
	   | [\u210A-\u2113]
	   | [\u2115]
	   | [\u2119-\u211D]
	   | [\u2124]
	   | [\u2126]
	   | [\u2128]
	   | [\u212A-\u212D]
	   | [\u212F-\u2131]
	   | [\u2133-\u2139]
	   | [\u2160-\u2183]
	   | [\u3005-\u3007]
	   | [\u3021-\u3029]
	   | [\u3031-\u3035]
	   | [\u3038-\u303A]
	   | [\u3041-\u3094]
	   | [\u309D-\u309E]
	   | [\u30A1-\u30FA]
	   | [\u30FC-\u30FE]
	   | [\u3105-\u312C]
	   | [\u3131-\u318E]
	   | [\u31A0-\u31B7]
	   | [\u3400]
	   | [\u4DB5]
	   | [\u4E00]
	   | [\u9FA5]
	   | [\uA000-\uA48C]
	   | [\uAC00]
	   | [\uD7A3]
	   | [\uF900-\uFA2D]
	   | [\uFB00-\uFB06]
	   | [\uFB13-\uFB17]
	   | [\uFB1D]
	   | [\uFB1F-\uFB28]
	   | [\uFB2A-\uFB36]
	   | [\uFB38-\uFB3C]
	   | [\uFB3E]
	   | [\uFB40-\uFB41]
	   | [\uFB43-\uFB44]
	   | [\uFB46-\uFBB1]
	   | [\uFBD3-\uFD3D]
	   | [\uFD50-\uFD8F]
	   | [\uFD92-\uFDC7]
	   | [\uFDF0-\uFDFB]
	   | [\uFE70-\uFE72]
	   | [\uFE74]
	   | [\uFE76-\uFEFC]
	   | [\uFF21-\uFF3A]
	   | [\uFF41-\uFF5A]
	   | [\uFF66-\uFFBE]
	   | [\uFFC2-\uFFC7]
	   | [\uFFCA-\uFFCF]
	   | [\uFFD2-\uFFD7]
	   | [\uFFDA-\uFFDC]
	*/
	return NewNfaWithRegular(func(c rune) bool {
		return ('\u0041' <= c && c <= '\u005A') ||
			('\u0061' <= c && c <= '\u007A') ||
			c == '\u00AA' ||
			c == '\u00B5' ||
			c == '\u00BA' ||
			('\u00C0' <= c && c <= '\u00D6') ||
			('\u00D8' <= c && c <= '\u00F6') ||
			('\u00F8' <= c && c <= '\u021F') ||
			('\u0222' <= c && c <= '\u0233') ||
			('\u0250' <= c && c <= '\u02AD') ||
			('\u02B0' <= c && c <= '\u02B8') ||
			('\u02BB' <= c && c <= '\u02C1') ||
			('\u02D0' <= c && c <= '\u02D1') ||
			('\u02E0' <= c && c <= '\u02E4') ||
			c == '\u02EE' ||
			c == '\u037A' ||
			c == '\u0386' ||
			('\u0388' <= c && c <= '\u038A') ||
			c == '\u038C' ||
			('\u038E' <= c && c <= '\u03A1') ||
			('\u03A3' <= c && c <= '\u03CE') ||
			('\u03D0' <= c && c <= '\u03D7') ||
			('\u03DA' <= c && c <= '\u03F3') ||
			('\u0400' <= c && c <= '\u0481') ||
			('\u048C' <= c && c <= '\u04C4') ||
			('\u04C7' <= c && c <= '\u04C8') ||
			('\u04CB' <= c && c <= '\u04CC') ||
			('\u04D0' <= c && c <= '\u04F5') ||
			('\u04F8' <= c && c <= '\u04F9') ||
			('\u0531' <= c && c <= '\u0556') ||
			c == '\u0559' ||
			('\u0561' <= c && c <= '\u0587') ||
			('\u05D0' <= c && c <= '\u05EA') ||
			('\u05F0' <= c && c <= '\u05F2') ||
			('\u0621' <= c && c <= '\u063A') ||
			('\u0640' <= c && c <= '\u064A') ||
			('\u0671' <= c && c <= '\u06D3') ||
			c == '\u06D5' ||
			('\u06E5' <= c && c <= '\u06E6') ||
			('\u06FA' <= c && c <= '\u06FC') ||
			c == '\u0710' ||
			('\u0712' <= c && c <= '\u072C') ||
			('\u0780' <= c && c <= '\u07A5') ||
			('\u0905' <= c && c <= '\u0939') ||
			c == '\u093D' ||
			c == '\u0950' ||
			('\u0958' <= c && c <= '\u0961') ||
			('\u0985' <= c && c <= '\u098C') ||
			('\u098F' <= c && c <= '\u0990') ||
			('\u0993' <= c && c <= '\u09A8') ||
			('\u09AA' <= c && c <= '\u09B0') ||
			c == '\u09B2' ||
			('\u09B6' <= c && c <= '\u09B9') ||
			('\u09DC' <= c && c <= '\u09DD') ||
			('\u09DF' <= c && c <= '\u09E1') ||
			('\u09F0' <= c && c <= '\u09F1') ||
			('\u0A05' <= c && c <= '\u0A0A') ||
			('\u0A0F' <= c && c <= '\u0A10') ||
			('\u0A13' <= c && c <= '\u0A28') ||
			('\u0A2A' <= c && c <= '\u0A30') ||
			('\u0A32' <= c && c <= '\u0A33') ||
			('\u0A35' <= c && c <= '\u0A36') ||
			('\u0A38' <= c && c <= '\u0A39') ||
			('\u0A59' <= c && c <= '\u0A5C') ||
			c == '\u0A5E' ||
			('\u0A72' <= c && c <= '\u0A74') ||
			('\u0A85' <= c && c <= '\u0A8B') ||
			c == '\u0A8D' ||
			('\u0A8F' <= c && c <= '\u0A91') ||
			('\u0A93' <= c && c <= '\u0AA8') ||
			('\u0AAA' <= c && c <= '\u0AB0') ||
			('\u0AB2' <= c && c <= '\u0AB3') ||
			('\u0AB5' <= c && c <= '\u0AB9') ||
			c == '\u0ABD' ||
			c == '\u0AD0' ||
			c == '\u0AE0' ||
			('\u0B05' <= c && c <= '\u0B0C') ||
			('\u0B0F' <= c && c <= '\u0B10') ||
			('\u0B13' <= c && c <= '\u0B28') ||
			('\u0B2A' <= c && c <= '\u0B30') ||
			('\u0B32' <= c && c <= '\u0B33') ||
			('\u0B36' <= c && c <= '\u0B39') ||
			c == '\u0B3D' ||
			('\u0B5C' <= c && c <= '\u0B5D') ||
			('\u0B5F' <= c && c <= '\u0B61') ||
			('\u0B85' <= c && c <= '\u0B8A') ||
			('\u0B8E' <= c && c <= '\u0B90') ||
			('\u0B92' <= c && c <= '\u0B95') ||
			('\u0B99' <= c && c <= '\u0B9A') ||
			c == '\u0B9C' ||
			('\u0B9E' <= c && c <= '\u0B9F') ||
			('\u0BA3' <= c && c <= '\u0BA4') ||
			('\u0BA8' <= c && c <= '\u0BAA') ||
			('\u0BAE' <= c && c <= '\u0BB5') ||
			('\u0BB7' <= c && c <= '\u0BB9') ||
			('\u0C05' <= c && c <= '\u0C0C') ||
			('\u0C0E' <= c && c <= '\u0C10') ||
			('\u0C12' <= c && c <= '\u0C28') ||
			('\u0C2A' <= c && c <= '\u0C33') ||
			('\u0C35' <= c && c <= '\u0C39') ||
			('\u0C60' <= c && c <= '\u0C61') ||
			('\u0C85' <= c && c <= '\u0C8C') ||
			('\u0C8E' <= c && c <= '\u0C90') ||
			('\u0C92' <= c && c <= '\u0CA8') ||
			('\u0CAA' <= c && c <= '\u0CB3') ||
			('\u0CB5' <= c && c <= '\u0CB9') ||
			c == '\u0CDE' ||
			('\u0CE0' <= c && c <= '\u0CE1') ||
			('\u0D05' <= c && c <= '\u0D0C') ||
			('\u0D0E' <= c && c <= '\u0D10') ||
			('\u0D12' <= c && c <= '\u0D28') ||
			('\u0D2A' <= c && c <= '\u0D39') ||
			('\u0D60' <= c && c <= '\u0D61') ||
			('\u0D85' <= c && c <= '\u0D96') ||
			('\u0D9A' <= c && c <= '\u0DB1') ||
			('\u0DB3' <= c && c <= '\u0DBB') ||
			c == '\u0DBD' ||
			('\u0DC0' <= c && c <= '\u0DC6') ||
			('\u0E01' <= c && c <= '\u0E30') ||
			('\u0E32' <= c && c <= '\u0E33') ||
			('\u0E40' <= c && c <= '\u0E46') ||
			('\u0E81' <= c && c <= '\u0E82') ||
			c == '\u0E84' ||
			('\u0E87' <= c && c <= '\u0E88') ||
			c == '\u0E8A' ||
			c == '\u0E8D' ||
			('\u0E94' <= c && c <= '\u0E97') ||
			('\u0E99' <= c && c <= '\u0E9F') ||
			('\u0EA1' <= c && c <= '\u0EA3') ||
			c == '\u0EA5' ||
			c == '\u0EA7' ||
			('\u0EAA' <= c && c <= '\u0EAB') ||
			('\u0EAD' <= c && c <= '\u0EB0') ||
			('\u0EB2' <= c && c <= '\u0EB3') ||
			('\u0EBD' <= c && c <= '\u0EC4') ||
			c == '\u0EC6' ||
			('\u0EDC' <= c && c <= '\u0EDD') ||
			c == '\u0F00' ||
			('\u0F40' <= c && c <= '\u0F6A') ||
			('\u0F88' <= c && c <= '\u0F8B') ||
			('\u1000' <= c && c <= '\u1021') ||
			('\u1023' <= c && c <= '\u1027') ||
			('\u1029' <= c && c <= '\u102A') ||
			('\u1050' <= c && c <= '\u1055') ||
			('\u10A0' <= c && c <= '\u10C5') ||
			('\u10D0' <= c && c <= '\u10F6') ||
			('\u1100' <= c && c <= '\u1159') ||
			('\u115F' <= c && c <= '\u11A2') ||
			('\u11A8' <= c && c <= '\u11F9') ||
			('\u1200' <= c && c <= '\u1206') ||
			('\u1208' <= c && c <= '\u1246') ||
			c == '\u1248' ||
			('\u124A' <= c && c <= '\u124D') ||
			('\u1250' <= c && c <= '\u1256') ||
			c == '\u1258' ||
			('\u125A' <= c && c <= '\u125D') ||
			('\u1260' <= c && c <= '\u1286') ||
			c == '\u1288' ||
			('\u128A' <= c && c <= '\u128D') ||
			('\u1290' <= c && c <= '\u12AE') ||
			c == '\u12B0' ||
			('\u12B2' <= c && c <= '\u12B5') ||
			('\u12B8' <= c && c <= '\u12BE') ||
			c == '\u12C0' ||
			('\u12C2' <= c && c <= '\u12C5') ||
			('\u12C8' <= c && c <= '\u12CE') ||
			('\u12D0' <= c && c <= '\u12D6') ||
			('\u12D8' <= c && c <= '\u12EE') ||
			('\u12F0' <= c && c <= '\u130E') ||
			c == '\u1310' ||
			('\u1312' <= c && c <= '\u1315') ||
			('\u1318' <= c && c <= '\u131E') ||
			('\u1320' <= c && c <= '\u1346') ||
			('\u1348' <= c && c <= '\u135A') ||
			('\u13A0' <= c && c <= '\u13B0') ||
			('\u13B1' <= c && c <= '\u13F4') ||
			('\u1401' <= c && c <= '\u1676') ||
			('\u1681' <= c && c <= '\u169A') ||
			('\u16A0' <= c && c <= '\u16EA') ||
			('\u1780' <= c && c <= '\u17B3') ||
			('\u1820' <= c && c <= '\u1877') ||
			('\u1880' <= c && c <= '\u18A8') ||
			('\u1E00' <= c && c <= '\u1E9B') ||
			('\u1EA0' <= c && c <= '\u1EE0') ||
			('\u1EE1' <= c && c <= '\u1EF9') ||
			('\u1F00' <= c && c <= '\u1F15') ||
			('\u1F18' <= c && c <= '\u1F1D') ||
			('\u1F20' <= c && c <= '\u1F39') ||
			('\u1F3A' <= c && c <= '\u1F45') ||
			('\u1F48' <= c && c <= '\u1F4D') ||
			('\u1F50' <= c && c <= '\u1F57') ||
			c == '\u1F59' ||
			c == '\u1F5B' ||
			c == '\u1F5D' ||
			('\u1F5F' <= c && c <= '\u1F7D') ||
			('\u1F80' <= c && c <= '\u1FB4') ||
			('\u1FB6' <= c && c <= '\u1FBC') ||
			c == '\u1FBE' ||
			('\u1FC2' <= c && c <= '\u1FC4') ||
			('\u1FC6' <= c && c <= '\u1FCC') ||
			('\u1FD0' <= c && c <= '\u1FD3') ||
			('\u1FD6' <= c && c <= '\u1FDB') ||
			('\u1FE0' <= c && c <= '\u1FEC') ||
			('\u1FF2' <= c && c <= '\u1FF4') ||
			('\u1FF6' <= c && c <= '\u1FFC') ||
			c == '\u207F' ||
			c == '\u2102' ||
			c == '\u2107' ||
			('\u210A' <= c && c <= '\u2113') ||
			c == '\u2115' ||
			('\u2119' <= c && c <= '\u211D') ||
			c == '\u2124' ||
			c == '\u2126' ||
			c == '\u2128' ||
			('\u212A' <= c && c <= '\u212D') ||
			('\u212F' <= c && c <= '\u2131') ||
			('\u2133' <= c && c <= '\u2139') ||
			('\u2160' <= c && c <= '\u2183') ||
			('\u3005' <= c && c <= '\u3007') ||
			('\u3021' <= c && c <= '\u3029') ||
			('\u3031' <= c && c <= '\u3035') ||
			('\u3038' <= c && c <= '\u303A') ||
			('\u3041' <= c && c <= '\u3094') ||
			('\u309D' <= c && c <= '\u309E') ||
			('\u30A1' <= c && c <= '\u30FA') ||
			('\u30FC' <= c && c <= '\u30FE') ||
			('\u3105' <= c && c <= '\u312C') ||
			('\u3131' <= c && c <= '\u318E') ||
			('\u31A0' <= c && c <= '\u31B7') ||
			c == '\u3400' ||
			c == '\u4DB5' ||
			c == '\u4E00' ||
			c == '\u9FA5' ||
			('\uA000' <= c && c <= '\uA48C') ||
			c == '\uAC00' ||
			c == '\uD7A3' ||
			('\uF900' <= c && c <= '\uFA2D') ||
			('\uFB00' <= c && c <= '\uFB06') ||
			('\uFB13' <= c && c <= '\uFB17') ||
			c == '\uFB1D' ||
			('\uFB1F' <= c && c <= '\uFB28') ||
			('\uFB2A' <= c && c <= '\uFB36') ||
			('\uFB38' <= c && c <= '\uFB3C') ||
			c == '\uFB3E' ||
			('\uFB40' <= c && c <= '\uFB41') ||
			('\uFB43' <= c && c <= '\uFB44') ||
			('\uFB46' <= c && c <= '\uFBB1') ||
			('\uFBD3' <= c && c <= '\uFD3D') ||
			('\uFD50' <= c && c <= '\uFD8F') ||
			('\uFD92' <= c && c <= '\uFDC7') ||
			('\uFDF0' <= c && c <= '\uFDFB') ||
			('\uFE70' <= c && c <= '\uFE72') ||
			c == '\uFE74' ||
			('\uFE76' <= c && c <= '\uFEFC') ||
			('\uFF21' <= c && c <= '\uFF3A') ||
			('\uFF41' <= c && c <= '\uFF5A') ||
			('\uFF66' <= c && c <= '\uFFBE') ||
			('\uFFC2' <= c && c <= '\uFFC7') ||
			('\uFFCA' <= c && c <= '\uFFCF') ||
			('\uFFD2' <= c && c <= '\uFFD7') ||
			('\uFFDA' <= c && c <= '\uFFDC')
	}, GoLexerNone)
}
