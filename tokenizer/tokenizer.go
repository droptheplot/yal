package tokenizer

import (
	"bufio"
	"bytes"
)

var literals []rune

func init() {
	literals = []rune{'%', '+', '-', '/', '*', '=', '>', '<', '.', '|', '!', '&', '[', ']'}

	for i := 'a'; i <= 'z'; i++ {
		literals = append(literals, i)
	}

	for i := 'A'; i <= 'Z'; i++ {
		literals = append(literals, i)
	}

	for i := '0'; i <= '9'; i++ {
		literals = append(literals, i)
	}
}

// Tokenize returns slice of tokens. For example, expression like:
//   (+ "hello" "world")
// Will be converted to:
//   "(", "+", "hello", "world", ")"
func Tokenize(src []byte) []string {
	var tokens []string
	var char rune
	var err error
	var inString bool
	var atom []rune

	reader := bufio.NewReader(bytes.NewReader(src))

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
	for i := range literals {
		if literals[i] == r {
			return true
		}
	}

	return false
}

func applyAtom(tokens *[]string, atom *[]rune) {
	if len(*atom) > 0 {
		*tokens = append(*tokens, string(*atom))
		*atom = (*atom)[:0]
	}
}
