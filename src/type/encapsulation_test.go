package _type

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	name string
	id string
	age int
}

func (e *Employee) String() string{
	fmt.Printf("in func String() e's address is %x\n", unsafe.Pointer(&e.name))
	return fmt.Sprintf("id:%s/name:%s/age:%d", e.id, e.name, e.age)
}

func TestStructOprations(t *testing.T){
	e := Employee{"张三", "1", 30}
	fmt.Printf("e's address is: %x\n", unsafe.Pointer(&e.name))
	t.Log(e.String())
	e2 := &Employee{"李四", "2", 25}
	fmt.Printf("e2's address is: %x\n", unsafe.Pointer(&e2.name))
	t.Log(e2.String())
}
