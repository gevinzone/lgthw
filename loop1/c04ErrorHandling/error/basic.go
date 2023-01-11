package main

import (
	"errors"
	"fmt"
)

// ValueError is a way to make a package level error to check against.
// I.e. if err == ErrorValue
var ValueError = errors.New("this is a value error")

// TypedError is a way to make an error type you can do err.(type) == ErrorValue
type TypedError struct {
	error
}

func DemoBasicError() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("err: ", err)

	DemoWrapError(err)

	err = ValueError
	fmt.Println("value error: ", err)

	var err2 error
	err2 = TypedError{err}
	target := TypedError{err}
	fmt.Println("TypedError{err} is a typed error, its ValueError value: ", err2)
	switch err2.(type) {
	case TypedError:
		fmt.Println("err2 is a typed error")
	default:
		fmt.Println("err2 is an unknown error")
	}
	fmt.Println("err2 is target? ", errors.Is(err2, target))
	fmt.Println("err2 == target? ", err2 == target)
}
