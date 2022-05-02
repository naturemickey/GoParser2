package lex

import "fmt"

type cursor struct {
	index  int
	line   int
	column int
}

func (this cursor) string() string {
	return fmt.Sprintf("at %d : %d", this.line, this.column)
}
