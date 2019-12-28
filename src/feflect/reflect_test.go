package feflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.ValueOf(f))
	t.Log(reflect.TypeOf(f))
}

func checkValueType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("float type")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int type")
	default:
		fmt.Println("unknown type", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 34
	checkValueType(f)
	checkValueType(&f)
}

type Employee struct {
	Age  int
	Name string `format:"normal"`
	Id   int
}

func (e *Employee) UpdateAge(newAge int) {
	e.Age = newAge
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{30, "张三", 11}
	t.Logf("Name: value(%[1]v), type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("failed get Name field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	//e.updateAge(20)
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(20)
	fmt.Println(in[0])
	// MethodByName 不支持仅包中可用的方法， 如 updateAge
	f := reflect.ValueOf(e).MethodByName("UpdateAge")
	fmt.Println(f)
	f.Call(in)

	t.Log("updated age", e)
}
