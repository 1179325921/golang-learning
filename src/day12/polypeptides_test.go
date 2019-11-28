package day12

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

func WriteHelloWorld(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.println(\"Hello world!!\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() string {
	return "System.out.println(\"Hello World!!\")"
}

func TestPoly(t *testing.T) {
	j := new(JavaProgrammer)
	g := new(GoProgrammer)
	WriteHelloWorld(j)
	WriteHelloWorld(g)
}
