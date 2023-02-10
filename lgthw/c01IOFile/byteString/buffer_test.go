package byteString

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuffer(t *testing.T) {
	var bufferTests = []struct {
		name  string
		input string
		want  *bytes.Buffer
	}{
		{
			name:  "normal",
			input: "abc",
			want:  bytes.NewBufferString("abc"),
		},
		{
			name:  "nil",
			input: "",
			want:  bytes.NewBufferString(""),
		},
		{
			name:  "special",
			input: "❤️",
			want:  bytes.NewBufferString("❤️"),
		},
	}

	for _, tt := range bufferTests {
		t.Run(tt.name, func(t *testing.T) {
			res := Buffer(tt.input)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestToString(t *testing.T) {
	var testCases = []struct {
		name    string
		input   *bytes.Buffer
		want    string
		wantErr error
	}{
		{
			name:    "normal",
			input:   bytes.NewBufferString("abc"),
			want:    "abc",
			wantErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res, err := toString(tt.input)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, res)
		})

	}
}
