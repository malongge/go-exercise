package unittest

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func square(num int) int {
	return num * num
}

func TestSquare(t *testing.T) {
	num1 := [...]int{1, 2, 3}
	num2 := [...]int{1, 4, 9}
	for i, value := range num1 {
		assert.Equal(t, num2[i], square(value))
	}

}

func TestStrAdd(t *testing.T) {
	a := assert.New(t)
	elems := [...]string{"1", "2", "3"}
	ret := ""
	for _, each := range elems {
		ret += each
	}
	a.Equal("123", ret)
}

func TestStrBufferAdd(t *testing.T) {
	var buf bytes.Buffer
	a := assert.New(t)
	elems := []string{"a", "b", "c"}
	for _, each := range elems {
		buf.WriteString(each)
	}
	t.Log(buf.Len())
	a.Equal("abc", buf.String())
}
