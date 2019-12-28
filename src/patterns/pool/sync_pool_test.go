package pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new object")
			return 10
		}}

	obj := pool.Get().(int)
	fmt.Println(obj)
	obj2 := pool.Get().(int)
	fmt.Println(obj2)
	pool.Put(9)
	obj, _ = pool.Get().(int)
	fmt.Println(obj)
	// GC 会清除sync.pool中缓存的对象
	runtime.GC()
	obj, _ = pool.Get().(int)
	fmt.Println(obj)
	pool.Put(3)
	obj, _ = pool.Get().(int)
	fmt.Println(obj)
	// 对象获取到之后，就会从 pool 中移除
	obj2, _ = pool.Get().(int)
	fmt.Println(obj2)

}

/*
sync.Pool 对象获取步骤
1.尝试从私有对象获取
2.私有对象不存在， 尝试从当前Processor的共享池获取
3.如果当前Processor共享池也是空的，那么就尝试去其他Processor的共享池获取
4.如果所有子池都是空的，最后就用用户指定的New函数产生一个新的对象返回

sync.Pool 对象放回步骤
1.如果私有对象不存在则保存为私有对象
2.如果私有对象存在，放入当前Processor子池的共享池中

sync.Pool 对象的生命周期
- GC 会清除 sync.pool 缓存的对象
- 对象的缓存有效期为下一次GC之前

sync.Pool 总结
- 适合于通过复用，降低复杂对象的创建和GC代价
- 协程安全， 会有锁的开销
- 生命周期受GC影响，不适合于做连接池等，需自己管理生命周期的资源的池化
*/
func TestSyncPoolInGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new object")
			return 10
		}}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)
	wg := new(sync.WaitGroup)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("%d: %d\n", i, pool.Get().(int))
			wg.Done()
		}(i)
	}
	wg.Wait()

}
