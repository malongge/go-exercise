package goroutine

import (
	"sync"
	"testing"
	"time"
)

// 共享内存没加保护
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter=%d", counter)
}


// 增加读写锁保护
func TestCounterThreadSafe(t *testing.T){
	var mut sync.RWMutex
	counter := 0
	for i:=0; i<5000; i++{
		go func() {
			mut.Lock()
			counter ++
			defer mut.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter=%d", counter)

}