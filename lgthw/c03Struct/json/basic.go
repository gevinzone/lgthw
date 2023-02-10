package json

import (
	"encoding/json"
	"fmt"
)

const (
	jsonBlob     = `{"name": "Aaron"}`
	fullJsonBlob = `{"name":"Aaron", "age":20}`
)

type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

func BaseEncoding() error {
	var e Example

	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular unmarshal: %+v\n", e)

	v, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular marshal, no age: ", string(v))

	if err = json.Unmarshal([]byte(fullJsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular unmarshal, with age: %+v\n", e)

	v, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular marshal, with age: ", string(v))

	return nil
}
