package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	fmt.Println("Hello Go")
	var a = []int{1, 2, 3, 4}
	// var b interface{}
	b := a[:3]
	fmt.Println(b)
	for _, v := range b { // Iterates over the arguments whatever the number.
		fmt.Println(v)
	}
	f, v := innerFunc(b...)
	data := f(30)
	fmt.Println(data)
	fmt.Println(v)
	convert()
	ss := SecondStruct{2}
	// ss.setN(10)
	ms := MyStruct{X:3, Y:4}
	// ms.myFunc()
	ms2 := MyStruct{3, 4, ss}
	// ms.setN(10)
	ms2.print()
	ms2.myFunc()
	ms2.xFunc()
	
	fmt.Println(ms.addWithP())
	fmt.Println(ms.addWithoutP())
	fmt.Println(ms.addWithoutP())

}

func innerFunc(args ...int) (func(val int) int, int) {
	fmt.Println(args)
	ret := 0
	for _, v := range args {
		// fmt.Println(v)
		ret = v + ret
		// fmt.Println(ret)
	}
	fmt.Println(ret)
	ret_func := func(val int) int {
		return ret + val + 50
	}

	return ret_func, ret
}

func convert() {
	i := 42
	fmt.Printf("%d\n", i)
	f := float64(i)
	fmt.Printf("%f\n", f)
	u := uint64(f)
	fmt.Printf("%d\n", u)

}

type SecondStruct struct {	

	N int
}


func (ss *SecondStruct) setN(value int){
	ss.N = value
}

func (ss *SecondStruct) print(){
	fmt.Printf("%d-----\n", ss.N)
}



type MyInterface interface{
	addWithP() int
	addWithoutP() int
	
	
}

type MyStruct struct{
	X, Y int
	// 合并时，指针和不是指针的区别是什么
	SecondStruct
}

func (mStruct *MyStruct) addWithP() int{
	mStruct.X += 5
	return mStruct.X + mStruct.Y
}

func (mStruct MyStruct) addWithoutP() int{
	mStruct.X += 5
	return mStruct.X + mStruct.Y
}

// func(mStruct *MyStruct) setN(value int) {
// 	mStruct.N = value
// }


func (mStruct MyStruct) myFunc(){
	mStruct.N = 10
	mStruct.print()
}

func (mStruct MyStruct) xFunc(){
	mStruct.print()
}
