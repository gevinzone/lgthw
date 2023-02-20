package reflection

import (
	"errors"
	"reflect"
)

var (
	errNotSupportedKind = errors.New("not supported kind")
	errReflectNil       = errors.New("nil is not supported")
	errReflectZero      = errors.New("zero is not supported")
)

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
