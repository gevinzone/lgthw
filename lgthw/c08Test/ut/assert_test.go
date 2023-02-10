package ut

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAssert(t *testing.T) {
	var err error = nil
	assert.NoError(t, err)
	expected, actual := 1, 1
	assert.Equal(t, expected, actual)
	assert.NotEqual(t, expected, actual+1)
	assert.Greater(t, actual+1, actual)
	assert.Less(t, actual, actual+1)

	// require 断言如果失败，下面代码不再继续执行
	// assert 断言如果失败，下面代码还会继续执行
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}
