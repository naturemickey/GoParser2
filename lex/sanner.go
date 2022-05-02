package lex

import "io/ioutil"

type scanner struct {
	code   []rune
	cursor cursor
}

func NewScannerFromFile(filePath string) *scanner {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}
	code := string(file)
	return NewScannerFromCode(code)
}

func NewScannerFromCode(code string) *scanner {
	return &scanner{code: []rune(code)}
}

func (this *scanner) LA() rune {
	char := this.code[this.cursor.index]
	this.cursor.index++
	return char
}

func (this *scanner) Next() bool {
	return this.cursor.index < len(this.code)
}

func (this *scanner) getTokenLiteral(from, to cursor) string {
	return string(this.code[from.index:to.index])
}
