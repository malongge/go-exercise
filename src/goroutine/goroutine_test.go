package goroutine

import (
	"fmt"
	"testing"
)

/* 输出结果是无序的
5
2
0
1
7
6
3
8
4
9
 */
func TestGoroutine(t *testing.T){
	for i:=0; i<10; i++{
		go func(i int){
			fmt.Println(i)
		}(i)
	}
}
