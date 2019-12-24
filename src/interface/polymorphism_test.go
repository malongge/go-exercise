package _interface

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {

	WriteHelloWorld() Code
}

type Go struct {


}

func (p *Go) WriteHelloWorld() Code{
	return "fmt.Println(\"hello world\")"
}

type Python struct {

}

func (p *Python) WriteHelloWorld() Code{
	return "print(\"hello world\")"
}

func writeFirstProgramm(p Programmer){
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestWriteFirstProgram(t *testing.T){
	g := new(Go)
	gaddr := &Go{}
	writeFirstProgramm(g)
	writeFirstProgramm(gaddr)
	p := new(Python)
	writeFirstProgramm(p)
	t.Log("program")
}