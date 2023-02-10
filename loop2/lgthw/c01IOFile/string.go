package c01IOFile

import (
	"bytes"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"strings"
)

func ModifyString() {
	s := "simple string"
	// prints [simple string]
	fmt.Println(strings.Split(s, " "))

	// prints "Simple String"
	fmt.Println(strings.Title(s))
	fmt.Println(cases.Title(language.Make("")).String(s))

	fmt.Println()
	fmt.Println(strings.ToTitle(s))
	fmt.Println(strings.ToUpper(s))

	// prints "simple string"; all trailing and
	// leading white space is removed
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

func StringToReader(s string) (io.Reader, error) {
	return strings.NewReader(s), nil
}

func StringToReaderV2(s string) (io.Reader, error) {
	return bytes.NewReader([]byte(s)), nil
}

func StringToReaderV3(s string) (io.Reader, error) {
	return bytes.NewBufferString(s), nil
}
