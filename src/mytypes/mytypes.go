package mytypes

import (
	"strings"
)

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyString string

func (s MyString) Len() int {
	return len(s)
}

type MyBuilder struct {
	Contents strings.Builder
}

type StringUppercaser struct {
	Contents strings.Builder
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

type MyBuilderEmbedded struct {
	strings.Builder
}
