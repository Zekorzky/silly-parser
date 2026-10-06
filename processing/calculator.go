package processing

import "strconv"

func Calculate(node *ASTNode) float64 {
	if node.Type == NUMBER {
		v, _ := strconv.ParseFloat(node.Value, 64)
		return v
	}

	var left float64
	var right float64

	if len(node.Children) == 1 {
		right = Calculate(node.Children[0])
	} else if len(node.Children) == 2 {
		left = Calculate(node.Children[0])
		right = Calculate(node.Children[1])
	}

	var result float64

	if node.Type == OPERATOR {
		result = twoInputTokenToFunction[node.Value](left, right)
	} else if node.Type == FUNCTION {
		result = oneInputTokenToFunction[node.Value](right)
	}

	return result
}
