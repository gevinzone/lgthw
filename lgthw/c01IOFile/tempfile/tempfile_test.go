package tempfile

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkWithTemp(t *testing.T) {
	err := WorkWithTemp()
	assert.NoError(t, err)
}

func TestWorkWithTemp1(t *testing.T) {
	err := WorkWithTemp1()
	assert.NoError(t, err)
}
