package feflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func fillBySettings(st interface{}, settings map[string]interface{}) error {
	//fmt.Println(reflect.TypeOf(st))
	//fmt.Println(reflect.TypeOf(st).Kind())
	errInfo := "the first param should be a pointer to struct type"
	if reflect.TypeOf(st).Kind() == reflect.Ptr {
		//fmt.Println(reflect.TypeOf(st).Elem())
		//fmt.Println(reflect.TypeOf(st).Elem().Kind())
		//fmt.Println(reflect.ValueOf(st).Elem())
		//fmt.Println(reflect.ValueOf(st).Elem().Type())
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New(errInfo)
		}
	} else {
		return errors.New(errInfo)
	}

	if settings == nil {
		return errors.New("settings should not be nil")
	}

	var field reflect.StructField
	var ok bool
	for k, v := range settings {
		if field, ok = reflect.TypeOf(st).Elem().FieldByName(k); !ok {
			continue
		}

		if field.Type == reflect.TypeOf(v) {
			vst := reflect.ValueOf(st)
			vstValue := vst.Elem()
			vstValue.FieldByName(k).Set(reflect.ValueOf(v))

		}

	}

	return nil
}

type Customer struct {
	Name string
	Some float32
	Id   int
}

func TestFillNameAndAge(t *testing.T) {
	e := &Employee{}
	settings := map[string]interface{}{"Name": "李四", "Age": 32, "Some": float32(1.2)}
	err := fillBySettings(e, settings)
	if err != nil {
		t.Error("Employee fill name and age failed")
	}
	fmt.Println(e)

	c := Customer{}

	err = fillBySettings(&c, settings)
	if err != nil {
		t.Error("Customer fill Name and Some failed")
	}
	fmt.Println(c)

}
