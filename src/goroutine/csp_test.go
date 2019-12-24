package goroutine

import (
	"fmt"
	"testing"
	"time"
)

/*
csp 并发模型，通过 channel 进行通信
*/

func Service() string {
	time.Sleep(time.Microsecond * 50)
	return "Service Done"
}

func OtherTask() {
	fmt.Println("work on other task")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("other task done")
}

func TestService(t *testing.T) {
	fmt.Println(Service())
	OtherTask()
}

func AsyncService() chan string{
	retCh := make(chan string, 1)
	go func() {
		ret := Service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exit")
	}()
	return retCh
}

func TestAsyncService(t *testing.T){
	// 由于异步服务时间较短， 一开始会执行 OtherTask , 遇到 sleep， 主线程阻塞了，这时候能够看到异步服务执行的过程
	retCh := AsyncService()
	OtherTask()
	fmt.Println(<-retCh)
}
