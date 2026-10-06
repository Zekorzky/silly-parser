package processing

import (
	"fmt"
	"math"
	"unicode"
)

var twoInputTokenToFunction map[string]func(a, b float64) float64 = map[string]func(a, b float64) float64{}
var oneInputTokenToFunction map[string]func(a float64) float64 = map[string]func(a float64) float64{}

func InitTokenizator() {

	twoInputTokenToFunction["+"] = func(a, b float64) float64 { return a + b }
	twoInputTokenToFunction["-"] = func(a, b float64) float64 { return a - b }
	twoInputTokenToFunction["*"] = func(a, b float64) float64 { return a * b }
	twoInputTokenToFunction["/"] = func(a, b float64) float64 { return a / b }
	twoInputTokenToFunction["^"] = func(a, b float64) float64 { return math.Pow(a, b) }

	oneInputTokenToFunction["sin"] = func(a float64) float64 { return math.Sin(a) }
	oneInputTokenToFunction["cos"] = func(a float64) float64 { return math.Cos(a) }
	oneInputTokenToFunction["tan"] = func(a float64) float64 { return math.Tan(a) }
	oneInputTokenToFunction["cot"] = func(a float64) float64 { return 1 / math.Tan(a) }
	oneInputTokenToFunction["ln"] = func(a float64) float64 { return math.Log(a) }
	oneInputTokenToFunction["abs"] = func(a float64) float64 { return math.Abs(a) }
	oneInputTokenToFunction["sqrt"] = func(a float64) float64 { return math.Sqrt(a) }
	oneInputTokenToFunction["deg"] = func(a float64) float64 { return a * math.Pi / 180 }
	oneInputTokenToFunction["fact"] = func(a float64) float64 { return Factorial(a) }
	oneInputTokenToFunction["gamma"] = func(a float64) float64 { return math.Gamma(a) }

}

func Factorial(a float64) float64 {
	// This project has too much recursion so yes.
	b := math.Round(a)
	if a != math.Trunc(a) {
		fmt.Println("Yk that factorials are for natural numbers only. We changed your", a, "to", b)
	}
	result := 1.0
	for b > 0 {
		result = result * b
		b--
	}
	return result
}

func AddToken(op string, fun func(a, b float64) float64) {
	twoInputTokenToFunction[op] = fun
}

func IsOperator(ch string) bool {
	return twoInputTokenToFunction[ch] != nil
}

func IsLetter(ch string) bool {
	return unicode.IsLetter(rune(ch[0]))
}

func IsToken(name string) bool {
	return oneInputTokenToFunction[name] != nil
}
