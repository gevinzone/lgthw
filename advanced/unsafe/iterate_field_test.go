package unsafe

import (
	"fmt"
	"reflect"
	"testing"
)

func printOffset(entity any) map[string]uintptr {
	typ := reflect.TypeOf(entity)
	fieldMap := make(map[string]uintptr, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		fd := typ.Field(i)
		fieldMap[fd.Name] = fd.Offset
		fmt.Println(fd.Name, fd.Offset)
	}
	return fieldMap
}

func TestPrintOffset(t *testing.T) {
	type user struct {
		name    string
		gender  byte
		gender2 int8
		gender3 int16
		gender4 int32
		age     int
		weight  int
		desc    string
		id      int64
	}
	_ = printOffset(user{})
}
