package encoding

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
)

func YamlEncoderDecoderExample() error {
	buffer := bytes.Buffer{}
	p := Pos{
		X:      0,
		Y:      0,
		Object: "something",
	}

	encoder := yaml.NewEncoder(&buffer)
	if err := encoder.Encode(p); err != nil {
		return err
	}
	fmt.Println("Yaml Encoded: ", string(buffer.Bytes()))

	var res Pos
	decoder := yaml.NewDecoder(&buffer)
	if err := decoder.Decode(&res); err != nil {
		return err
	}
	fmt.Println("Yaml decoded: ", res)
	fmt.Println("Yaml decoded equals origin: ", res == p)
	return nil
}

func YamlMarshalUnMarshalExample() error {
	p := Pos{
		X:      0,
		Y:      0,
		Object: "something",
	}

	b, err := yaml.Marshal(&p)
	if err != nil {
		return err
	}
	fmt.Println("Yaml Encoded: ", string(b))

	var res Pos
	err = yaml.Unmarshal(b, &res)
	if err != nil {
		return err
	}
	fmt.Println("Yaml decoded: ", res)
	fmt.Println("Yaml decoded equals origin: ", res == p)
	return nil
}
