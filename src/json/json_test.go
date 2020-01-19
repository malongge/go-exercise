package json

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var jsonStr = `{
	"info": {
		"name": "张三",
		"age": 30
	},
	"job": {
		"skills": ["开发", "测试"]
	}
}`

func TestJson(t *testing.T) {
	a := assert.New(t)
	e := new(Employee)
	if err := json.Unmarshal([]byte(jsonStr), e); err != nil {
		a.Error(errors.New("json decoded error"))
	}
	a.Equal(e.BasicInfo.Name, "张三")
	t.Log(*e)
	v, err := json.Marshal(e)
	if err != nil {
		a.Error(errors.New("json encoded error"))
	}
	t.Log(string(v))
}

func TestEasyJson(t *testing.T) {
	e := Employee{}
	a := assert.New(t)
	if err := e.UnmarshalJSON([]byte(jsonStr)); err != nil {
		a.Error(errors.New("easy json decoded error"))
	}
	a.Equal(e.BasicInfo.Name, "张三")
	t.Log(e)
	v, err := e.MarshalJSON()
	if err != nil {
		a.Error(errors.New("easy json encode error"))
	}
	t.Log(string(v))
}

func BenchmarkJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal([]byte(jsonStr), e); err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		if err := e.UnmarshalJSON([]byte(jsonStr)); err != nil {
			b.Error(err)
		}
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}
