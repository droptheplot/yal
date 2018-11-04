package tokenizer_test

import (
	"fmt"
	"testing"

	"github.com/droptheplot/yal/tokenizer"
	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	var testCases = []struct {
		src    []byte
		result []string
	}{
		{
			[]byte("(qwe asd zxc)"),
			[]string{"(", "qwe", "asd", "zxc", ")"},
		},
		{
			[]byte("(qwe (asd zxc))"),
			[]string{"(", "qwe", "(", "asd", "zxc", ")", ")"},
		},
		{
			[]byte("(qwe (asd (zxc)))"),
			[]string{"(", "qwe", "(", "asd", "(", "zxc", ")", ")", ")"},
		},
		{
			[]byte("(+ 1 2)"),
			[]string{"(", "+", "1", "2", ")"},
		},
		{
			[]byte("(qwe \"asd zxc\")"),
			[]string{"(", "qwe", "\"asd zxc\"", ")"},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.src), func(t *testing.T) {
			assert.Equal(t, tc.result, tokenizer.Tokenizer{}.Tokenize(tc.src))
		})
	}
}
