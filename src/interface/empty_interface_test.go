package _interface

import (
	"fmt"
	"testing"
)

func judgeType(p interface{}){

	switch v:=p.(type){
	case int:
		fmt.Println("integer: ", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknow type")
	}
}

func TestJudegeType(t *testing.T){
	judgeType(1)
	judgeType("1")
	judgeType(make([]int, 5))
}