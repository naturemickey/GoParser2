package lex

import "fmt"

type cursor struct {
	index  int
	line   int
	column int
}

func newCursor() cursor {
	return cursor{index: 0, line: 1, column: 0}
}

func (this cursor) string() string {
	return fmt.Sprintf("at %d : %d", this.line, this.column)
}
