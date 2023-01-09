package json

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// NullInt64 实现了 json.Marshaler 和 json.Unmarshaler 接口
type NullInt64 sql.NullInt64

type ExampleNullInt struct {
	Age  *NullInt64 `json:"age,omitempty"`
	Name string
}

func (n *NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Int64)
}

func (n *NullInt64) UnmarshalJSON(b []byte) error {
	n.Valid = false
	if b == nil {
		return nil
	}
	n.Valid = true
	return json.Unmarshal(b, &n.Int64)
}

var _ json.Marshaler = &NullInt64{}
var _ json.Unmarshaler = &NullInt64{}

func EncodingNull() error {
	var e ExampleNullInt
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("nullInt64 Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("nullInt64 Marshal, with no age:", string(value))

	if err = json.Unmarshal([]byte(fullJsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("nullInt64 Unmarshal, with age = 20: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("nullInt64 Marshal, with age = 20:", string(value))

	return nil
}
