package reflection

import "reflect"

func IterateFunc(entity any) (map[string]FuncInfo, error) {
	if entity == nil {
		return nil, errReflectNil
	}
	typ := reflect.TypeOf(entity)
	val := reflect.ValueOf(entity)
	if val.IsZero() {
		return nil, errReflectZero
	}

	numMethod := typ.NumMethod()
	res := make(map[string]FuncInfo)
	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)
		fn := method.Func

		numIn := fn.Type().NumIn()
		inputTypes := make([]reflect.Type, 0, numIn)
		inputValues := make([]reflect.Value, 0, numIn)

		inputTypes = append(inputTypes, reflect.TypeOf(entity))
		inputValues = append(inputValues, reflect.ValueOf(entity))
		for j := 1; j < numIn; j++ {
			t := fn.Type().In(j)
			inputTypes = append(inputTypes, t)
			inputValues = append(inputValues, reflect.Zero(t))
		}

		numOut := fn.Type().NumOut()
		outputTypes := make([]reflect.Type, numOut)
		for j := 0; j < numOut; j++ {
			outputTypes[j] = fn.Type().Out(j)
		}
		outputs := fn.Call(inputValues)
		result := make([]any, len(outputs))
		for j, out := range outputs {
			result[j] = out.Interface()
		}

		res[method.Name] = FuncInfo{
			Name:        method.Name,
			InputTypes:  inputTypes,
			OutputTypes: outputTypes,
			Result:      result,
		}
	}
	return res, nil
}

type FuncInfo struct {
	Name        string
	InputTypes  []reflect.Type
	OutputTypes []reflect.Type
	Result      []any
}
