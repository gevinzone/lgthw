package filedir

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Capitalize opens a file, reads and changes the contents, then writes them to a second file
func Capitalize(f1, f2 *os.File) error {
	if _, err := f1.Seek(0, 0); err != nil {
		return err
	}
	var tmp = new(bytes.Buffer)
	if _, err := io.Copy(tmp, f1); err != nil {
		return err
	}
	s := strings.ToUpper(tmp.String())
	if _, err := io.Copy(f2, strings.NewReader(s)); err != nil {
		return err
	}
	return nil
}

func CapitalizeExample() error {
	fileName1, fileName2 := "file1.txt", "file2.txt"
	var (
		f1  *os.File
		f2  *os.File
		err error
	)
	data := []byte(`
    this file contains
    a number of words
    and new lines`)
	if f1, err = os.Create(fileName1); err != nil {
		return err
	}
	if _, err = f1.Write(data); err != nil {
		return err
	}
	if f2, err = os.Create(fileName2); err != nil {
		return err
	}

	if err = Capitalize(f1, f2); err != nil {
		return err
	}

	var res []byte
	if res, err = ioutil.ReadFile(fileName2); err != nil {
		return err
	}
	fmt.Println("res: ", string(res))

	if err = removeFiles(fileName1, fileName2); err != nil {
		return err
	}
	// 提前close的话，f2就读不到数据了
	//_ = f2.Close()
	// 如果没有这一句，f2读不到数据
	_, _ = f2.Seek(0, 0)
	res, _ = ioutil.ReadAll(f2)
	fmt.Println("f2: ", string(res))
	_ = f1.Close()
	_ = f2.Close()
	return nil
}

func removeFiles(files ...string) error {
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return err
		}
	}
	return nil
}
