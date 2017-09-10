package main

import (
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
		"int64":  regexp.MustCompile(`^[0-9]+$`),
		"bool":   regexp.MustCompile(`^(true|false)$`),
		"string": regexp.MustCompile(`^\".+\"$`),
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
	src = strings.Replace(src, "(", " ( ", -1)
	src = strings.Replace(src, ")", " ) ", -1)

	var tokens []string

	for _, token := range strings.Split(src, " ") {
		if token != "" && token != "\n" && token != "\n\n" {
			tokens = append(tokens, token)
		}
	}

	return tokens
}

func parse(tokens []string) (Node, int) {
	if Debug {
		fmt.Println(tokens)
	}

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
		_value, _ = strconv.ParseInt(node.atom, 10, 32)
		return _value
	case "bool":
		var _value bool
		_value, _ = strconv.ParseBool(node.atom)
		return _value
	case "string":
		var _value string
		_value, _ = strconv.Unquote(node.atom)
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
