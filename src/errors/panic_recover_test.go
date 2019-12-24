package errors

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestPanicWithoutExit(t *testing.T){
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("recover from err: ", err)
		}
	}()

	panic(errors.New("something wrong happened"))
}
