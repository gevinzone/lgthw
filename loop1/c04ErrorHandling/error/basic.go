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

	// %w 用于wrap error
	err1 := fmt.Errorf("some wrapped error: %w", err)
	fmt.Println("err1: ", err1)
	fmt.Println("err1 is err? ", errors.Is(err1, err))
	err0 := errors.Unwrap(err1)
	fmt.Println("unwrap err1 to err0:", err0)
	fmt.Println("err0 == err? ", errors.Is(err0, err), err0 == err)

	err = ValueError
	fmt.Println("value error: ", err)

	var err2 error
	err2 = TypedError{err}
	target := TypedError{err}
	fmt.Println("typed error: ", err2)
	switch err2.(type) {
	case TypedError:
		fmt.Println("typed error: ", err2)
	default:
		fmt.Println("unknown error: ", err2)
	}
	fmt.Println("err2 is target? ", errors.Is(err2, target))
}
