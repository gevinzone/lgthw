package unsafe

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccessor_Field(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	u := &user{
		name: "Tom",
		age:  18,
	}
	accessor := NewAccessor(u)
	name, err := accessor.Field("name")
	require.NoError(t, err)
	assert.Equal(t, "Tom", name)
	age, err := accessor.Field("age")
	require.NoError(t, err)
	assert.Equal(t, 18, age)

	err = accessor.SetField("age", 19)
	require.NoError(t, err)
	err = accessor.SetField("name", "Jerry")
	require.NoError(t, err)
	assert.Equal(t, "Jerry", u.name)
	assert.Equal(t, 19, u.age)
	assert.Equal(t, &user{
		name: "Jerry",
		age:  19,
	}, u)
}
