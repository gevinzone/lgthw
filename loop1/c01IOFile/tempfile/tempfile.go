package tempfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WorkWithTemp will give some basic patterns for working
// with temporary files and directories
func WorkWithTemp() error {
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}

	defer os.RemoveAll(t)

	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())
	if _, err = tf.Write([]byte("tmp data")); err != nil {
		return err
	}
	return tf.Close()
}

func WorkWithTemp1() error {
	content := []byte("temporary file's content")
	tmpDir, err := ioutil.TempDir("./", "example")
	if err != nil {
		return err
	}
	fmt.Println(tmpDir)
	defer os.RemoveAll(tmpDir)
	tmpFile := filepath.Join(tmpDir, "tmpFile")
	if err = ioutil.WriteFile(tmpFile, content, 0666); err != nil {
		return err
	}
	return nil
}
