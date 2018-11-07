package parser_test

import (
	"fmt"
	"testing"

	"github.com/droptheplot/yal/parser"
	"github.com/droptheplot/yal/yal"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var testCases = []struct {
		tokens []string
		result yal.Node
	}{
		{
			[]string{"(", "qwe", ")"},
			yal.Node{Nodes: []yal.Node{
				yal.Node{Atom: "qwe"},
			}},
		},
		{
			[]string{"(", "qwe", "asd", "zxc", ")"},
			yal.Node{Nodes: []yal.Node{
				yal.Node{Atom: "qwe", Nodes: []yal.Node{
					yal.Node{Atom: "asd"},
					yal.Node{Atom: "zxc"},
				}},
			}},
		},
		{
			[]string{"(", "qwe", "(", "asd", "zxc", ")", ")"},
			yal.Node{Nodes: []yal.Node{
				yal.Node{Atom: "qwe", Nodes: []yal.Node{
					yal.Node{Atom: "asd", Nodes: []yal.Node{
						yal.Node{Atom: "zxc"},
					}},
				}},
			}},
		},
		{
			[]string{"(", "qwe", "(", "asd", "(", "zxc", ")", ")", ")"},
			yal.Node{Nodes: []yal.Node{
				yal.Node{Atom: "qwe", Nodes: []yal.Node{
					yal.Node{Atom: "asd", Nodes: []yal.Node{
						yal.Node{Atom: "zxc"},
					}},
				}},
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.tokens), func(t *testing.T) {
			assert.Equal(t, tc.result, parser.Parser{}.Parse(tc.tokens))
		})
	}
}
