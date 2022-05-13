package util

import (
	"go/format"
)

func GoFmt(content string) string {
	bs, err := format.Source([]byte(content))
	if err != nil {
		panic(err)
	}
	return string(bs)
}
