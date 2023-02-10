package main

import "github.com/gevinzone/lgthw/lgthw/c03Struct/encoding"

func main() {
	encoding.Base64Example()
	_ = encoding.Base64ExampleEncoderDecoder()
	_ = encoding.GobExample()
	_ = encoding.JsonEncoderDecoderExample()
	_ = encoding.JsonMarshalUnMarshalExample()
	_ = encoding.YamlEncoderDecoderExample()
	_ = encoding.YamlMarshalUnMarshalExample()
}
