package c01IOFile

import "io"

func Copy(in io.Reader, out io.Writer) (int64, error) {
	return io.Copy(out, in)
}

func CopyV2(in io.Reader, out io.Writer) (int64, error) {
	buf := make([]byte, 64)
	return io.CopyBuffer(out, in, buf)
}
