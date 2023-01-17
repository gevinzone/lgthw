package c01IOFile

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBuffer(t *testing.T) {
	s := "it's easy to encode unicode into a byte array ❤️"
	bs, err := WriteStringToBuffer(s)
	assert.NoError(t, err)

	bs2, err := WriteStringToBufferV2(s)
	assert.NoError(t, err)

	bs3, err := WriteStringToBufferV3(s)
	assert.NoError(t, err)

	assert.Equal(t, bs, bs3)
	assert.Equal(t, bs2, bs3)

	s2, err := ReadStringFromBuffer(bs)
	assert.NoError(t, err)
	assert.Equal(t, s, s2)

	assert.Equal(t, 0, bs.Len())
}

func TestString(t *testing.T) {
	ModifyString()
	s := "This is a test"
	r1, _ := StringToReader(s)
	r2, _ := StringToReaderV2(s)
	r3, _ := StringToReaderV3(s)
	s1, _ := getFromReader(r1)
	s2, _ := getFromReader(r2)
	s3, _ := getFromReader(r3)
	assert.Equal(t, s1, s2)
	assert.Equal(t, s1, s3)

}

func getFromReader(r io.Reader) (string, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func TestCopy(t *testing.T) {
	s := "This is a testing string\n"
	r := strings.NewReader(s)
	written, err := Copy(r, os.Stdout)
	assert.NoError(t, err)
	assert.Equal(t, r.Size(), written)

	_, err = r.Seek(0, 0)
	assert.NoError(t, err)
	written, err = CopyV2(r, os.Stdout)
	assert.NoError(t, err)
	assert.Equal(t, r.Size(), written)
}
