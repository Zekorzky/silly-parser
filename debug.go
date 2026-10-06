package main

import (
	"fmt"
	"parser-test/processing"
)

func DebugTree(ast *processing.ASTNode) {
	printOutTree(ast, "", true, 0)
}

func printOutTree(head *processing.ASTNode, prefix string, isLast bool, depth int) {
	if depth == 0 {
		fmt.Println(head.Value)
		for i, c := range head.Children {
			printOutTree(c, "", i == len(head.Children)-1, depth+1)
		}
		return
	}

	connector := "├── "
	if isLast {
		connector = "└── "
	}
	fmt.Println(prefix + connector + head.Value)

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	for i, c := range head.Children {
		printOutTree(c, childPrefix, i == len(head.Children)-1, depth+1)
	}
}
