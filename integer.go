package main

func SUM(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]).(int64), eval(node.nodes[1]).(int64)

	return a + b
}

func SUB(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]).(int64), eval(node.nodes[1]).(int64)

	return a - b
}

func MUL(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]).(int64), eval(node.nodes[1]).(int64)

	return a * b
}

func DIV(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]).(int64), eval(node.nodes[1]).(int64)

	return a / b
}

func MOD(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]).(int64), eval(node.nodes[1]).(int64)

	return a % b
}

func ISGT(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]).(int64), eval(node.nodes[1]).(int64)

	return a > b
}

func ISLT(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]).(int64), eval(node.nodes[1]).(int64)

	return a < b
}

func ISGTEQ(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]).(int64), eval(node.nodes[1]).(int64)

	return a >= b
}

func ISLTEQ(node Node) interface{} {
	check_arity(node, 2)

	a, b := eval(node.nodes[0]).(int64), eval(node.nodes[1]).(int64)

	return a <= b
}
