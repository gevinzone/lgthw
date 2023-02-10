package filedir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperate(t *testing.T) {
	err := Operate()
	assert.NoError(t, err)
}
