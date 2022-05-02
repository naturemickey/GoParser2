package lex

import "io/ioutil"

type scanner struct {
	code   []rune
	index  int
	line   int
	column int
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
	return this.code[this.index]
}

func (this *scanner) Next() bool {
	this.index++
	return this.index < len(this.code)
}
