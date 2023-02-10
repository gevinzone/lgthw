package encoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Pos struct {
	X      int
	Y      int
	Object string
}

func GobExample() error {
	buffer := bytes.Buffer{}
	p := Pos{
		X:      0,
		Y:      0,
		Object: "something",
	}

	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(p); err != nil {
		return err
	}
	fmt.Println("Gob Encoded: ", buffer.Bytes())

	var res Pos
	decoder := gob.NewDecoder(&buffer)
	if err := decoder.Decode(&res); err != nil {
		return err
	}
	fmt.Println("Gob decoded: ", res)
	fmt.Println("Gob decoded equals origin: ", res == p)
	return nil
}
