package goroutine

import (
	"fmt"
	"testing"
)

func sum(a []int, c chan int) {
	sum_ := 0
	for _, v := range a {
		sum_ += v
	}
	c <- sum_
	fmt.Println(sum_)
}

func TestSum(t *testing.T) {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)
	go sum(a, c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z, x+y+z)

}
