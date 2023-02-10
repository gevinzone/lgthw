package main

import "github.com/gevinzone/lgthw/loop1/c03Struct/encoding"

func main() {
	encoding.Base64Example()
	_ = encoding.Base64ExampleEncoderDecoder()
	_ = encoding.GobExample()
	_ = encoding.JsonEncoderDecoderExample()
	_ = encoding.JsonMarshalUnMarshalExample()
	_ = encoding.YamlEncoderDecoderExample()
	_ = encoding.YamlMarshalUnMarshalExample()
}
