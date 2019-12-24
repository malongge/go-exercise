package errors

import (
	"errors"
	"fmt"
	"testing"
)

var NegtiveError = errors.New("负数错误")
var LargeThanHundredError = errors.New("超过100错误")

func score(num int)(int, error) {
	if num < 0 {
		return num, NegtiveError
	} else if num > 100 {
		return num, LargeThanHundredError
	}
	fmt.Println("the number is:", num)
	return num, nil
}

func TestScore(t *testing.T){
	_, err := score(-1)
	if err == NegtiveError{
		fmt.Println("test pass")
	}else{
		t.Error(err)
	}
	_, err = score(101)
	if err == LargeThanHundredError{
		fmt.Println("test pass")
	}else{
		t.Error(err)
	}
}
