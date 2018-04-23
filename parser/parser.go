package parser

import "github.com/droptheplot/yal/yal"

// Parse returns root node for given tokens.
// Last return variable indicates shift inside current subtree and should be ommited by user.
func Parse(tokens []string) (yal.Node, int) {
	var node = yal.Node{}
	var i int

	for i < len(tokens) {
		token := tokens[i]

		if token == "(" {
			tmp, move := Parse(tokens[i+1:])
			node.Nodes = append(node.Nodes, tmp)
			i += move + 1
		} else if token == ")" {
			return node, i
		} else if node.Atom == "" {
			node.Atom = token
		} else {
			var valueNode = yal.Node{Atom: token}
			node.Nodes = append(node.Nodes, valueNode)
		}

		i++
	}

	return node, len(tokens)
}
