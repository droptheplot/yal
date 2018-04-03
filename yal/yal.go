package yal

import (
	"bufio"
	"strings"
)

type Node struct {
	Atom  string
	Nodes []Node
}

func New(src []byte) Node {
	tokens := tokenize(string(src))
	node, _ := parse(tokens)

	return node
}

func parse(tokens []string) (Node, int) {
	var token string
	var size int = len(tokens)
	var i int = 0

	node := Node{}

	for i < size {
		token = tokens[i]

		if token == "(" {
			tmp, move := parse(tokens[i+1:])
			node.Nodes = append(node.Nodes, tmp)
			i += move + 1
		} else if token == ")" {
			return node, i
		} else if node.Atom == "" {
			node.Atom = token
		} else {
			var valueNode = Node{Atom: token}
			node.Nodes = append(node.Nodes, valueNode)
		}

		i++
	}

	return node, size
}

func tokenize(src string) []string {
	var tokens []string
	var char rune
	var err error = nil
	var inString bool = false
	var atom []rune

	reader := bufio.NewReader(strings.NewReader(src))

	for err == nil {
		char, _, err = reader.ReadRune()

		if inString {
			if isQuote(char) {
				inString = false
				atom = append(atom, char)
				tokens = append(tokens, string(atom))
				atom = atom[:0]
			} else {
				atom = append(atom, char)
			}

			continue
		}

		switch {
		case isLParen(char):
			tokens = append(tokens, string(char))
		case isRParen(char):
			applyAtom(&tokens, &atom)
			tokens = append(tokens, string(char))
		case isLiteral(char):
			atom = append(atom, char)
		case isQuote(char):
			inString = true
			atom = append(atom, char)
		case isWhitespace(char):
			applyAtom(&tokens, &atom)
		}
	}

	return tokens
}

func isLParen(r rune) bool {
	return r == '('
}

func isRParen(r rune) bool {
	return r == ')'
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n'
}

func isQuote(r rune) bool {
	return r == '"'
}

func isLiteral(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') ||
		r == '%' || r == '+' || r == '-' || r == '/' || r == '*' ||
		r == '=' || r == '>' || r == '<' || r == '.'
}

func applyAtom(tokens *[]string, atom *[]rune) {
	if len(*atom) > 0 {
		*tokens = append(*tokens, string(*atom))
		*atom = (*atom)[:0]
	}
}
