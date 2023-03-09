package reflection

import (
	"errors"
	"reflect"
)

var (
	errNotSupportedKind = errors.New("not supported kind")
	errReflectNil       = errors.New("nil is not supported")
	errReflectZero      = errors.New("zero is not supported")
	errFieldCanSet      = errors.New("field can not be set")
)

// IterateFields 把结构体的字段，展开存储到map中，也支持操作指向结构体的指针
// 如果结构体中的字段组合了其他结构体或指针，不再递归展开
func IterateFields(entity any) (map[string]any, error) {
	if entity == nil {
		return nil, errReflectNil
	}
	typ := reflect.TypeOf(entity)
	val := reflect.ValueOf(entity)

	if val.IsZero() {
		return nil, errReflectZero
	}

	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, errNotSupportedKind
	}

	numField := typ.NumField()
	res := make(map[string]any)
	for i := 0; i < numField; i++ {
		fieldType := typ.Field(i)
		fieldVal := val.Field(i)
		if fieldType.IsExported() {
			res[fieldType.Name] = fieldVal.Interface()
		} else {
			res[fieldType.Name] = reflect.Zero(fieldType.Type).Interface()
		}
	}
	return res, nil
}

func SetField(entity any, fieldName string, newVal any) error {
	if entity == nil {
		return errReflectNil
	}
	val := reflect.ValueOf(entity)
	if val.IsZero() {
		return errReflectZero
	}
	for val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return errNotSupportedKind
	}
	fieldVal := val.FieldByName(fieldName)
	if fieldVal.IsZero() || !fieldVal.CanSet() {
		return errFieldCanSet
	}
	fieldVal.Set(reflect.ValueOf(newVal))
	return nil
}
