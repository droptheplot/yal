package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/kr/pretty"
)

var Env map[string]interface{}
var ClassRules map[string]*regexp.Regexp
var Debug bool
var File string

func init() {
	Env = map[string]interface{}{
		"+":      SUM,
		"-":      SUB,
		"*":      MUL,
		"/":      DIV,
		"%":      MOD,
		"==":     ISEQ,
		">":      ISGT,
		">=":     ISGTEQ,
		"<":      ISLT,
		"<=":     ISLTEQ,
		"if":     IF,
		"print":  PRINT,
		"def":    DEF,
		"repeat": REPEAT,
	}

	ClassRules = map[string]*regexp.Regexp{
		"int64": regexp.MustCompile(`^[0-9]+$`),
		"bool":  regexp.MustCompile(`^(true|false)$`),
	}
}

func main() {
	flag.StringVar(&File, "file", "", "Path to file")
	flag.BoolVar(&Debug, "debug", false, "Debug mode")

	flag.Parse()

	if File == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	src, err := ioutil.ReadFile(File)

	check_error(err)

	tokens := tokenize(string(src))
	ast, _ := parse(tokens)

	if Debug {
		fmt.Printf("%# v\n\n", pretty.Formatter(ast))
	}

	result := eval(ast)

	fmt.Println(result)
}

type Node struct {
	atom  string
	nodes []Node
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
			if isQUOTE(char) {
				inString = false
				tokens = append(tokens, string(atom))
				atom = atom[:0]
			} else {
				atom = append(atom, char)
			}

			continue
		}

		switch {
		case isLPAR(char):
			tokens = append(tokens, string(char))
		case isRPAR(char):
			applyAtom(&tokens, &atom)
			tokens = append(tokens, string(char))
		case isLITERAL(char):
			atom = append(atom, char)
		case isINTEGER(char):
			atom = append(atom, char)
		case isQUOTE(char):
			inString = true
		case isWHITESPACE(char):
			applyAtom(&tokens, &atom)
		}
	}

	return tokens
}

func isLPAR(r rune) bool {
	return r == '('
}

func isRPAR(r rune) bool {
	return r == ')'
}

func isWHITESPACE(r rune) bool {
	return r == ' ' || r == '\n'
}

func isQUOTE(r rune) bool {
	return r == '"'
}

func isINTEGER(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLITERAL(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '=' || r == '%'
}

func applyAtom(tokens *[]string, atom *[]rune) {
	if len(*atom) > 0 {
		*tokens = append(*tokens, string(*atom))
		*atom = (*atom)[:0]
	}
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
			node.nodes = append(node.nodes, tmp)
			i += move + 1
		} else if token == ")" {
			return node, i
		} else if inEnv(token) {
			node.atom = token
		} else {
			var valueNode = Node{atom: token}
			node.nodes = append(node.nodes, valueNode)
		}

		i++
	}

	return node, size
}

func eval(node Node) interface{} {
	if node.atom == "" {
		var result interface{}

		for _, n := range node.nodes {
			result = eval(n)
		}

		return result
	} else {
		if inEnv(node.atom) {
			if reflect.TypeOf(Env[node.atom]).String() == "main.Node" {
				return eval(Env[node.atom].(Node))
			} else {
				return Env[node.atom].(func(node Node) interface{})(node)
			}
		} else {
			return parseAtom(node)
		}
	}
}

func parseAtom(node Node) interface{} {
	var class string

	for k, regexp := range ClassRules {
		if regexp.MatchString(node.atom) {
			class = k
			break
		}
	}

	switch class {
	case "int64":
		var _value int64
		_value, _ = strconv.ParseInt(node.atom, 10, 64)
		return _value
	case "bool":
		var _value bool
		_value, _ = strconv.ParseBool(node.atom)
		return _value
	default:
		return node.atom
	}
}

func inEnv(atom string) bool {
	_, ok := Env[atom]

	return ok
}

func check_arity(node Node, arity int) {
	if len(node.nodes) != arity {
		panic("Wrong number of arguments")
	}
}

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}
