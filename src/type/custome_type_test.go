package _type

import (
	"fmt"
	"testing"
	"time"
)

type FuncType func(op int) int

func timeSpent(inner FuncType) FuncType {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFuncType(t *testing.T) {
	tsFunc := timeSpent(slowFunc)
	t.Log("return: ", tsFunc(10))
}
