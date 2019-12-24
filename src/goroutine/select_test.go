package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Microsecond * 500)
	return "Service Done"
}

func asyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	retCh := asyncService()
	select {
	case ret := <-retCh:
		t.Log(ret)
	case <-time.After(time.Microsecond * 1000):
		t.Error("service time out")

	}
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			fmt.Println("在通道中放入数据", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println(<-c)
			return
		}
	}
}


func TestFibonacci(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func(){
		for i:=0; i<10; i++{
			fmt.Println("当前迭代 i:", i)
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	time.Sleep(1*time.Second)
}
