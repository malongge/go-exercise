package test_const

import (
	"strconv"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wensday
)

const (
	R = 1 << iota
	W
	X
)

func TestArray(t *testing.T) {

	a1 := [...]int{1, 2, 3, 4}
	a2 := [...]int{2, 3, 4, 5}
	a3 := [...]int{1, 2, 3, 4}
	t.Log(a1 == a2)
	t.Log(a1 == a3)
}

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wensday)
	rwx := 7
	notRead := rwx &^ R
	t.Log(R, W, X)
	t.Log(R&notRead == R, W&notRead == W, X&notRead == X)
}

func TestSlice(t *testing.T) {
	s1 := []int{}
	s2 := []int{2, 3, 4, 5}
	n := 0
	for n < 5 {
		s1 = append(s1, n)
		n++
		t.Log(len(s1), cap(s1))
	}

	t.Log(s2)

	s2 = append(s2, 6, 7, 8, 9, 10)
	s3 := s2[4:6]
	t.Log(s3)
	s4 := s2[5:7]
	t.Log(s4)
	s4[0] = 88
	t.Log(s3)
	t.Log(s4)
	t.Log(len(s4), cap(s4))

	s5 := make([]int, 5, 10)
	t.Log(len(s5), cap(s5))

}

type myInt uint64

func TestType(t *testing.T) {
	i := 64
	// 不支持隐示类型转化
	//t.Log(myInt(i) == i)
	t.Log(myInt(i))

	var j int32 = 55
	t.Log(int(j) == i)
}

func TestStrings(t *testing.T) {
	str := "10"
	if val, err := strconv.Atoi(str); err == nil {
		t.Logf("type: %[1]T, value: %[1]d", val)
	}
	i := 10
	t.Logf("type: %[1]T, value: %[1]s", strconv.Itoa(i))

	v := "hello world"
	t.Log("bytes length is: ", len(v))
	v2 := "中"
	t.Log("bytes length is: ", len(v2))
	t.Logf("utf-8 store: %x", v2)
	c := []rune(v2)
	t.Log("unicode length is: ", len(c))
	t.Logf("unicode encode is: %x", c[0])
}

func TestMap(t *testing.T) {
	m := map[int]int{1: 2, 3: 4}
	t.Log(m[1])
	for k, v := range m {
		t.Log(k, v)
	}

	t.Log("m[10] value is default value: ", m[10])

	m2 := make(map[string]func(int) int)
	m2["a"] = func(v int) int { return v }
	m2["b"] = func(v int) int { return v * v }
	m2["c"] = func(v int) int { return v * v * v }

	for k, v := range m2 {
		t.Log(k, v(2))
	}


}

func TestSet(t *testing.T) {
	s := map[int]bool{}
	s[2] = true
	s[1] = true
	t.Log(len(s))
	if _, ok := s[1]; ok {
		t.Log("1 is exist in set")
	}
	delete(s, 1)
	if _, ok := s[1]; !ok {
		t.Log("1 is not exist in set")
	}
	// 借助默认值，可不用if判断
	t.Log(s[3])

}
