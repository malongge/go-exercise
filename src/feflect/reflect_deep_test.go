package feflect

import (
	"reflect"
	"testing"
)

type Things struct {
	Id int
}

func TestDeepEqual(t *testing.T) {
	var1 := map[int]string{1: "你好", 2: "我好", 3: "他好"}
	var2 := map[int]string{1: "你好", 2: "我好", 3: "他好"}
	// 无法直接用 == 比较两个字典是否相等，可以用反射的 DeepEqual
	t.Log(reflect.DeepEqual(var1, var2))
	list1 := [3]int{1, 2, 3}
	list2 := [3]int{1, 2, 3}
	// 像列表是可以比较的，但分片也不行
	t.Log(list1 == list2)
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{3, 2, 1}
	t.Log(reflect.DeepEqual(slice1, slice2))
	t.Log(reflect.DeepEqual(slice1, slice3))

	thing1 := &Things{1}
	thing2 := &Things{1}
	//
	t.Log(thing1 == thing2)
	t.Log(reflect.DeepEqual(thing1, thing2))
}
