package test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Excutable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Friday)
	a := 1 //0001
	t.Log(a&Readable == Readable)
}
