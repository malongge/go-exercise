package pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReuseableObj struct {
}

type Pool struct {
	objs chan *ReuseableObj
}

func newPool(num int) *Pool {
	objs := make(chan *ReuseableObj, num)
	pool := Pool{objs}
	for i := 0; i < num; i++ {
		obj := new(ReuseableObj)
		pool.objs <- obj
	}
	return &pool
}

func (pool *Pool) getObj(timeout time.Duration) (*ReuseableObj, error) {
	select {
	case obj := <-pool.objs:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("get object time out")
	}
}

func (pool *Pool) putObj(obj *ReuseableObj) error {
	select {
	case pool.objs <- obj:
		return nil
	default:
		return errors.New("pool is full")
	}
}

func TestNewPool(t *testing.T) {
	pool := newPool(10)
	if len(pool.objs) != 10 {
		t.Error("new pool failed")
	}
}

func TestGetObj(t *testing.T) {
	pool := newPool(1)
	_, err := pool.getObj(1 * time.Second)
	if err != nil {
		t.Error("from pool get object error")
	}
	obj, _ := pool.getObj(1 * time.Second)
	if obj != nil {
		t.Error("should not get object from empty pool")
	}

}

func TestPutObj(t *testing.T) {
	pool := newPool(2)
	err := pool.putObj(&ReuseableObj{})

	if err == nil {
		t.Error("full pool can not insert object")
	}
	_, err = pool.getObj(1 * time.Second)
	err = pool.putObj(&ReuseableObj{})
	if err != nil {
		t.Error("not full pool can be insert object")
	}

}

func TestObjPool(t *testing.T) {
	pool := newPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.getObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.putObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Printf("the pool size remain: %d\n", len(pool.objs))
	fmt.Println("Done")
}
