package day11

import (
	"testing"
)

type Programmer interface {
	WriteCode() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteCode() string {
	return "go语言开发写代码"
}

func TestWriteCode(t *testing.T) {
	g := GoProgrammer{}
	t.Log(g.WriteCode())
}
