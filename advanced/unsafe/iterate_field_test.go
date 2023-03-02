package unsafe

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// printOffset 打印结构体的字段和偏移量，参数类型支撑结构体和指向结构体的多级指针
func printOffset(entity any) (map[string]uintptr, error) {
	typ := reflect.TypeOf(entity)
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("not supported kind")
	}
	fieldMap := make(map[string]uintptr, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		fd := typ.Field(i)
		fieldMap[fd.Name] = fd.Offset
		fmt.Println(fd.Name, fd.Offset)
	}
	return fieldMap, nil
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
	t.Log("user{}:")
	_, _ = printOffset(user{})
	t.Log("&user{}:")
	_, _ = printOffset(&user{})
}
