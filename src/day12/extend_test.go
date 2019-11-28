package day12

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(host)
}

type Dog struct {
	//p *Pet
	Pet //匿名嵌套类型
}

func (d *Dog) Speak() {
	fmt.Println("汪汪汪.")
}

//func (d *Dog) SpeakTo(host string) {
//	//d.Speak()
//	//fmt.Println(host)
//	d.p.SpeakTo(host)
//}

func TestExtend(t *testing.T) {
	d := new(Dog)
	d.SpeakTo("狗")
}
