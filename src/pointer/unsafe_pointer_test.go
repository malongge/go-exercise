package pointer

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type MyInt int

func TestUnsafePointer(t *testing.T) {
	val := 6
	val2 := int64(val)
	t.Logf("va12:%[1]T, val2's value: %[1]d", val2)

	val3 := float32(val2)
	fmt.Println(val3)

	val4 := float64(7)
	val5 := float32(val4)
	fmt.Println(val4, val5)
	ptr := (*float64)(unsafe.Pointer(&val))
	fmt.Println(&val)
	fmt.Println(ptr)
	// 输出并非是 6， 非安全的转换
	fmt.Println(*ptr)
	f := *ptr
	fmt.Println(reflect.TypeOf(f))

}

// 类型别名转换是安全的
func TestMyINt(t *testing.T) {
	var1 := []int{1, 2, 3, 4}
	var2 := *(*[]MyInt)(unsafe.Pointer(&var1))
	fmt.Println(var2)
}

func TestAtomic(t *testing.T) {
	var sharePtr unsafe.Pointer
	writeDataFunc := func() {
		data := []int{}
		for i := 0; i < 10; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&sharePtr, unsafe.Pointer(&data))
	}

	readDataFunc := func() {
		data := atomic.LoadPointer(&sharePtr)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFunc()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFunc()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)

		go func() {
			for i := 0; i < 10; i++ {
				readDataFunc()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
	}
	readDataFunc()
	wg.Wait()
}
