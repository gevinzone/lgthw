package reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIterateFields(t *testing.T) {
	testCases := []struct {
		name    string
		entity  any
		wantVal map[string]any
		wantErr error
	}{
		{
			name: "struct with full fields",
			entity: User{
				Name:     "Tom",
				age:      18,
				Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
			},
			wantVal: map[string]any{
				"Name":     "Tom",
				"age":      0,
				"Birthday": time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "struct with nil fields",
			entity: User{
				age: 18,
			},
			wantVal: map[string]any{
				"Name":     "",
				"age":      0,
				"Birthday": time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "pointer",
			entity: &User{
				Name:     "Tom",
				age:      18,
				Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
			},
			wantVal: map[string]any{
				"Name":     "Tom",
				"age":      0,
				"Birthday": time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "pointer of zero",
			entity: func() *User {
				var u User
				return &u
			}(),
			wantVal: map[string]any{
				"Birthday": time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				"Name":     "",
				"age":      0,
			},
		},
		{
			name: "multi pointer",
			entity: func() **User {
				u := &User{
					Name:     "Tom",
					age:      18,
					Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
				}
				return &u
			}(),
			wantVal: map[string]any{
				"Name":     "Tom",
				"age":      0,
				"Birthday": time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name:    "nil",
			entity:  nil,
			wantErr: errReflectNil,
		},
		{
			name:    "nil with type",
			entity:  (*User)(nil),
			wantErr: errReflectZero,
		},
		{
			name:    "basic",
			entity:  18,
			wantErr: errNotSupportedKind,
		},
		{
			name:    "func",
			entity:  func() {},
			wantErr: errNotSupportedKind,
		},
		{
			name:    "slice",
			entity:  []int{1, 2, 3},
			wantErr: errNotSupportedKind,
		},
		{
			name:    "array",
			entity:  [...]int{1, 2, 3},
			wantErr: errNotSupportedKind,
		},
		{
			name: "map",
			entity: map[string]any{
				"k1": 1,
				"k2": "value",
			},
			wantErr: errNotSupportedKind,
		},
		{
			name: "composition in struct",
			entity: Person{
				Id: 1,
				User: User{
					Name:     "Tom",
					age:      18,
					Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
			wantVal: map[string]any{
				"Id": 1,
				"User": User{
					Name:     "Tom",
					age:      18,
					Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
		},
		{
			name: "composite pointer in struct",
			entity: Actor{
				Id: 1,
				User: &User{
					Name:     "Tom",
					age:      18,
					Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
			wantVal: map[string]any{
				"Id": 1,
				"User": &User{
					Name:     "Tom",
					age:      18,
					Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFields(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, res)
		})
	}
}

func TestSetField(t *testing.T) {
	testCases := []struct {
		name string

		entity any
		field  string
		newVal any

		wantEntity any
		wantErr    error
	}{
		{
			// 结构体存储在栈中，是unaddressable的
			name: "struct",
			entity: User{
				Name: "Tom",
			},
			field:   "Name",
			newVal:  "Jerry",
			wantErr: errFieldCanSet,
		},

		{
			name: "pointer",
			entity: &User{
				Name: "Tom",
			},
			field:  "Name",
			newVal: "Jerry",
			wantEntity: &User{
				Name: "Jerry",
			},
		},

		{
			name: "pointer not exported",
			entity: &User{
				age: 18,
			},
			field:   "age",
			newVal:  19,
			wantErr: errFieldCanSet,
		},
		{
			name:    "nil",
			entity:  nil,
			wantErr: errReflectNil,
		},
		{
			name:    "nil with type",
			entity:  (*User)(nil),
			wantErr: errReflectZero,
		},
		{
			name:    "not struct",
			entity:  18,
			wantErr: errNotSupportedKind,
		},
		{
			name:    "not struct",
			entity:  []int{1, 2, 3},
			wantErr: errNotSupportedKind,
		},
		{
			name:    "point to not struct",
			entity:  &([]int{1, 2, 3}),
			wantErr: errNotSupportedKind,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SetField(tc.entity, tc.field, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantEntity, tc.entity)
		})
	}
}
