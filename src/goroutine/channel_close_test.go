package goroutine

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func productor(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

/*
获取11次值的输出
0
1
2
3
4
5
6
7
8
9
0
加通道关闭判断后的输出结果
0
1
2
3
4
5
6
7
8
9
*/
func consumer(ch chan int, wg *sync.WaitGroup) {

	go func() {
		for i := 0; i < 11; i++ {
			// 通道中没有值时会返回零值
			//fmt.Println(<-ch)
			// 通过判断通道是否关闭，来结束处理
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()

}

func TestCloseChannel(t *testing.T) {

	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	productor(ch, &wg)
	wg.Add(1)
	consumer(ch, &wg)
	wg.Wait()
}

func isCanceled(cc chan struct{}) bool {
	select {
	case <-cc:
		return true
	default:
		return false
	}
}

func canceled1(cc chan struct{}) {
	cc <- struct{}{}
}

// close 会关闭掉所有 goroutine 中的通道
func canceled2(cc chan struct{}) {
	close(cc)
}

func TestCancel(t *testing.T) {
	var cc = make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cc chan struct{}) {
			for {
				if isCanceled(cc) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i)
		}(i, cc)

	}
	// 只会 cancel 掉一个线程
	canceled1(cc)
	// 会 cancel 掉所有的线程
	canceled2(cc)
	// 等待线程执行完成
	time.Sleep(1 * time.Second)
}

func isCanceledCtx(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

/*
根Context：通过context.Background()创建
子Context：context.WithCancel(parentContext)创建
ctx, cancel := context.WithCancel(context.Background())
当前 Context 被取消时，基于他的子 context 都会被取消
接收取消通知 <-ctx.Done()
 */
func TestCancelCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceledCtx(ctx) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i)
		}(i, ctx)
	}

	cancel()
	time.Sleep(1 * time.Second)
}
