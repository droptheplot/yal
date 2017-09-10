package main

import (
	"fmt"
	"strconv"
)

func ISEQ(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]), eval(node.nodes[1])

	return a == b
}

func IF(node Node) interface{} {
	var condition bool = eval(node.nodes[0]).(bool)

	if condition {
		return eval(node.nodes[1])
	} else if len(node.nodes) == 3 {
		return eval(node.nodes[2])
	}

	return nil
}

func PRINT(node Node) interface{} {
	check_arity(node, 1)

	a := eval(node.nodes[0])

	fmt.Println(a)

	return nil
}

func DEF(node Node) interface{} {
	check_arity(node, 2)

	name := node.nodes[0].atom
	value := node.nodes[1]

	Env[name] = value

	return nil
}

func REPEAT(node Node) interface{} {
	check_arity(node, 3)

	name := parseAtom(node.nodes[0]).(string)
	n := parseAtom(node.nodes[1]).(int64)

	for i := int64(0); i < n; i++ {
		Env[name] = Node{atom: strconv.FormatInt(i, 10)}
		eval(node.nodes[2])
	}

	delete(Env, name)

	return nil
}
