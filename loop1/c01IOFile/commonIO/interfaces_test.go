package commonIO

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name    string
		in      io.ReadSeeker
		want    io.Writer
		wantErr error
	}{
		{
			name:    "base-case",
			in:      bytes.NewReader([]byte("example")),
			want:    bytes.NewBufferString("exampleexample"),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := Copy(tt.in, out)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, out)
		})
	}
}

func TestCopyV1(t *testing.T) {
	tests := []struct {
		name    string
		in      io.Reader
		want    string
		wantErr error
	}{
		{
			name:    "base-case",
			in:      bytes.NewReader([]byte("example")),
			want:    "example",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := CopyV1(tt.in, out)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, out.String())
		})
	}
}

func TestCopyV2(t *testing.T) {
	tests := []struct {
		name    string
		in      io.Reader
		want    string
		wantErr error
	}{
		{
			name:    "base-case",
			in:      bytes.NewReader([]byte("this is an example")),
			want:    "this is an example",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := CopyV2(tt.in, out)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, out.String())
		})
	}
}
