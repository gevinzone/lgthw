package filedir

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Operate() error {
	// same with dir := "./example"
	dir := "example"
	//dir := "/tmp/example"
	// this will create a directory /tmp/example
	// you may also use an absolute path instead of a relative one
	if err := os.Mkdir(dir, os.FileMode(0755)); err != nil {
		return err
	}

	fmt.Println(filepath.Abs(dir))
	printCurrentWorkDir()

	// go to the ./example directory
	if err := os.Chdir(dir); err != nil {
		return err
	}
	printCurrentWorkDir()

	fileName := "test.txt"
	// f is a generic file object
	// it also implements multiple interfaces and can be used as a reader or writer
	// if the correct bits are set when opening
	f, err := os.Create(fileName)
	fmt.Println(f.Name())
	if err != nil {
		return err
	}
	data := []byte("hello\n")
	if n, err := f.Write(data); err != nil {
		return err
	} else {
		if n != len(data) {
			return errors.New("incorrect length returned from write")
		}

	}
	if err = f.Close(); err != nil {
		return err
	}

	f, err = os.Open(fileName)
	fmt.Println(f.Name())
	_, _ = io.Copy(os.Stdout, f)
	_ = f.Close()

	if err = os.Chdir(".."); err != nil {
		return err
	}
	printCurrentWorkDir()

	if err = os.RemoveAll(dir); err != nil {
		return err
	}
	return nil
}

func printCurrentWorkDir() {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println("current work dir: ", path)
}
