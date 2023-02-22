package reflection

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestIterateFunc(t *testing.T) {
	p := createUser()

	testCases := []struct {
		name    string
		entity  any
		wantRes map[string]FuncInfo
		wantErr error
	}{
		{
			name: "struct",
			entity: User{
				Name:     "Tom",
				age:      18,
				Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
			},
			wantRes: map[string]FuncInfo{
				"GetAge": {
					Name:        "GetAge",
					InputTypes:  []reflect.Type{reflect.TypeOf(User{})},
					OutputTypes: []reflect.Type{reflect.TypeOf(0)},
					Result:      []any{18},
				},
			},
		},
		{
			name:   "pointer",
			entity: p,
			wantRes: map[string]FuncInfo{
				"GetAge": {
					Name:        "GetAge",
					InputTypes:  []reflect.Type{reflect.TypeOf(&User{})},
					OutputTypes: []reflect.Type{reflect.TypeOf(0)},
					Result:      []any{18},
				},
				"ChangeName": {
					Name:        "ChangeName",
					InputTypes:  []reflect.Type{reflect.TypeOf(&User{}), reflect.TypeOf("")},
					OutputTypes: []reflect.Type{reflect.TypeOf(&User{})},
					Result:      []any{p},
				},
			},
		},
		{
			name:    "nil",
			entity:  nil,
			wantErr: errReflectNil,
		},
		{
			name:    "typed nil",
			entity:  (*User)(nil),
			wantErr: errReflectZero,
		},
		{
			name:    "int",
			entity:  18,
			wantRes: map[string]FuncInfo{},
		},
		{
			name:    "slice",
			entity:  []int{1, 2, 3},
			wantRes: map[string]FuncInfo{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFunc(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func createUser() *User {
	return &User{
		Name:     "Tom",
		age:      18,
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
}
