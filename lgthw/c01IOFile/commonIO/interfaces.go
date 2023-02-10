package commonIO

import (
	"fmt"
	"io"
	"os"
)

// Copy copies data from in to out first directly,
// then using a buffer. It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	_, _ = in.Seek(0, 0)
	buf := make([]byte, 1)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}
	fmt.Println()
	return nil
}

// The CopyV1 Copy() function copies between interfaces and treats them like streams.
// Thinking of data as streams has a lot of practical uses,
// especially when working with network traffic or filesystems
func CopyV1(in io.Reader, out io.Writer) error {
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}

func CopyV2(in io.Reader, out io.Writer) error {
	buf := make([]byte, 1)
	if _, err := io.CopyBuffer(out, in, buf); err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}

// PipeExample The primary purpose of a pipe is to read from a stream
// while simultaneously writing from the same stream to a different source.
// In essence, it combines the two streams into a pipe
func PipeExample() error {
	r, w := io.Pipe()
	go func() {
		_, _ = w.Write([]byte("this is an example"))
		_ = w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
