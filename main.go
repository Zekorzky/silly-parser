package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"parser-test/processing"
)

func main() {

	debug := flag.Bool("debug", false, "a bool")
	flag.Parse()

	processing.InitTokenizator()

	for {
		fmt.Printf("Enter your expression: ")
		var input string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()

		cleaned := processing.Clean(input)
		ast := processing.CreateAST(cleaned)

		if *debug {
			DebugTree(ast)
		}

		fmt.Println(processing.Calculate(ast))
	}
}
