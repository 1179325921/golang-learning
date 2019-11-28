package day10

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

//方法被调用时，实例的成员会进行值复制
func (e Employee) work() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id %s Name %s Age %d正在工作", e.Id, e.Name, e.Age)
}

//这个方法可以避免值复制
func (e *Employee) work1() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id %s Name %s Age %d正在工作", e.Id, e.Name, e.Age)
}

func TestEmployee(t *testing.T) {
	e1 := Employee{"1", "小明", 20}
	e2 := Employee{Name: "小红", Age: 25}
	e3 := new(Employee) //返回的是指针
	e3.Age = 30
	e3.Name = "张三"
	e3.Id = "3"
	t.Logf("e1 is %T", e1)
	t.Logf("e2 is %T", e2)
	t.Logf("e3 is %T", e3)
}

func TestEmployeeWork(t *testing.T) {
	e1 := Employee{"1", "小明", 20}
	t.Logf("Address is %x\n", unsafe.Pointer(&e1.Name))
	t.Log(e1.work())
	t.Log(e1.work1())
}
