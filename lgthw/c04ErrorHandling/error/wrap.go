package main

import (
	"errors"
	"fmt"
)

func DemoWrapError(err error) {
	// %w 用于wrap error
	err1 := fmt.Errorf("some wrapped error: %w", err)
	fmt.Println("err1: ", err1)
	fmt.Println("err1 is err? ", errors.Is(err1, err))
	err0 := errors.Unwrap(err1)
	fmt.Println("unwrap err1 to err0:", err0)
	fmt.Println("err0 == err? ", errors.Is(err0, err), err0 == err)
}
