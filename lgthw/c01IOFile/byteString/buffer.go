package byteString

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer demonstrates some tricks for initializing bytes Buffers
// These buffers implement an io.Reader interface
func Buffer(rawString string) *bytes.Buffer {
	b := &bytes.Buffer{}
	buf := []byte(rawString)
	b.Write(buf)
	//return b

	//alternative 1
	return bytes.NewBuffer(buf)
	//
	// alternative 2
	//return bytes.NewBufferString(rawString)

}

// ToString is an example of taking an io.Reader and consuming
// it all, then returning a string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), err
}
