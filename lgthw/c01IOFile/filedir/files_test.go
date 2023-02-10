package filedir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapitalizeExample(t *testing.T) {
	err := CapitalizeExample()
	assert.NoError(t, err)
}
