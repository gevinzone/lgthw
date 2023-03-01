package unsafe

import (
	"errors"
	"reflect"
	"unsafe"
)

type Accessor struct {
	fields  map[string]FieldMeta
	address unsafe.Pointer
}

func NewAccessor(entity any) *Accessor {
	typ := reflect.TypeOf(entity)
	typ = typ.Elem()

	fields := make(map[string]FieldMeta, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		fd := typ.Field(i)
		fields[fd.Name] = FieldMeta{
			Offset: fd.Offset,
			typ:    fd.Type,
		}
	}

	val := reflect.ValueOf(entity)
	return &Accessor{
		fields:  fields,
		address: val.UnsafePointer(),
	}
}

func (a *Accessor) Field(field string) (any, error) {
	meta, ok := a.fields[field]
	if !ok {
		return nil, errors.New("不合法字段")
	}
	typ := meta.typ
	p := unsafe.Pointer(uintptr(a.address) + meta.Offset)
	val := reflect.NewAt(typ, p)
	return val.Elem().Interface(), nil
}

func (a *Accessor) SetField(field string, val any) error {
	meta, ok := a.fields[field]
	if !ok {
		return errors.New("不合法字段")
	}
	p := unsafe.Pointer(uintptr(a.address) + meta.Offset)
	v := reflect.NewAt(meta.typ, p).Elem()
	v.Set(reflect.ValueOf(val))
	return nil
}

type FieldMeta struct {
	Offset uintptr
	typ    reflect.Type
}
