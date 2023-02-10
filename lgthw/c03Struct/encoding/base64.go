package encoding

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func Base64Example() {
	var value string
	s := "https://blog.igevin.info/post/a-b#a?x=%c"
	value = base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("With EncodeToString and StdEncoding", value)
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return
	}
	fmt.Println("With DecodeString and StdEncoding: ", string(decoded))

	value = base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println("With EncodeToString and urlEncoding", value)
	decoded, err = base64.URLEncoding.DecodeString(value)
	if err != nil {
		return
	}
	fmt.Println("With DecodeString and URLEncoding: ", string(decoded))
}

func Base64ExampleEncoderDecoder() error {
	buffer := bytes.Buffer{}
	s := "some data"

	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)
	if err := encoder.Close(); err != nil {
		return err
	}
	if _, err := encoder.Write([]byte(s)); err != nil {
		return err
	}
	fmt.Println("Using encoder and StdEncoding: ", buffer.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)
	result, err := ioutil.ReadAll(decoder)
	if err != nil {
		return err
	}
	fmt.Println("Using decoder and StdEncoding: ", string(result))
	fmt.Println("Decoded equals origin? ", string(result) == s)
	return nil
}
