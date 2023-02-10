package c01IOFile

import (
	"bytes"
	"io"
	"io/ioutil"
)

func WriteStringToBuffer(s string) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte(s))
	return buf, nil
}
func WriteStringToBufferV2(s string) (*bytes.Buffer, error) {
	return bytes.NewBufferString(s), nil
}
func WriteStringToBufferV3(s string) (*bytes.Buffer, error) {
	return bytes.NewBuffer([]byte(s)), nil
}

func ReadStringFromBuffer(r io.Reader) (string, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
