package extend

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	s := "这是golang的中文字符串"
	t.Log("字节长度: ", len(s))
	t.Log("字符串长度1: ", len([]rune(s)))
	t.Log("字符串长度2: ", utf8.RuneCountInString(s))
	t.Log("字符串长度3: ", utf8.RuneCount([]byte(s)))
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name    string
		val     string
		wantVal string
		wantErr error
	}{
		{
			name:    "ascii odd",
			val:     "abc1234567890",
			wantVal: "0987654321cba",
		},
		{
			name:    "ascii even",
			val:     "abcd1234567890",
			wantVal: "0987654321dcba",
		},
		{
			name:    "unicode odd",
			val:     "中文奇数个字符",
			wantVal: "符字个数奇文中",
		},
		{
			name:    "unicode even",
			val:     "中文偶数字符",
			wantVal: "符字数偶文中",
		},
		{
			name:    "unicode + ascii odd",
			val:     "中文a奇b数个字符",
			wantVal: "符字个数b奇a文中",
		},
		{
			name:    "unicode + ascii even",
			val:     "中文a偶数b字符",
			wantVal: "符字b数偶a文中",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ReverseString(tc.val)
			assert.Equal(t, tc.wantVal, res)
		})
	}

}
