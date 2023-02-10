package encoding

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonEncoderDecoderExample() error {
	buffer := bytes.Buffer{}
	p := Pos{
		X:      0,
		Y:      0,
		Object: "something",
	}

	encoder := json.NewEncoder(&buffer)
	if err := encoder.Encode(p); err != nil {
		return err
	}
	fmt.Println("Json Encoded: ", string(buffer.Bytes()))

	var res Pos
	decoder := json.NewDecoder(&buffer)
	if err := decoder.Decode(&res); err != nil {
		return err
	}
	fmt.Println("Json decoded: ", res)
	fmt.Println("Json decoded equals origin: ", res == p)
	return nil
}

func JsonMarshalUnMarshalExample() error {
	p := Pos{
		X:      0,
		Y:      0,
		Object: "something",
	}

	b, err := json.Marshal(&p)
	if err != nil {
		return err
	}
	fmt.Println("Json Encoded: ", string(b))

	var res Pos
	err = json.Unmarshal(b, &res)
	if err != nil {
		return err
	}
	fmt.Println("Json decoded: ", res)
	fmt.Println("Json decoded equals origin: ", res == p)
	return nil
}
