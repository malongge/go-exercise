package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result from %d", id)
}

func firstResponse() string {
	number := 10
	// 这里只会有一个协程产生， 任意任务完成就返回了
	//ch := make(chan string, number)
	// 利用 channel 的机制，防止协程泄露需要使用buffer channel, 这种情况下会产生 10 个协程
	ch := make(chan string)
	for i := 0; i < number; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	return <-ch
}

func allResponse() string{
	number := 10
	//ch := make(chan string)
	ch := make(chan string, number)
	for i:=0; i<number; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	final := "\n"
	for j:=0; j<number; j++ {
		temp := <- ch
		final += temp + "\n"
	}
	return final
}

func TestFirstResponse(t *testing.T) {
	t.Log("before run testing, goroutine number: ", runtime.NumGoroutine())
	t.Log(firstResponse())
	time.Sleep(time.Second * 1)
	t.Log("after run testing, goroutine number: ", runtime.NumGoroutine())
}

func TestAllResponse(t *testing.T){
	t.Log("before run testing, goroutine number: ", runtime.NumGoroutine())
	t.Log(allResponse())
	time.Sleep(time.Second * 1)
	t.Log("after run testing, goroutine number: ", runtime.NumGoroutine())
}
