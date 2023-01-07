package main

import "fmt"

type CustomError struct {
	msg string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("error message: %s", c.msg)
}

func DemoCustomError() error {
	return &CustomError{msg: "my fault"}
}
