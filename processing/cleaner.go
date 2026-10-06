package processing

import (
	"strconv"
)

func Clean(input string) []string {
	result := []string{}
	interval := ""
	for _, ch := range input {
		if IsOperator(string(ch)) || string(ch) == "(" || string(ch) == ")" {
			if interval != "" {
				println(interval)
				result = append(result, interval)
				interval = ""
			}
			result = append(result, string(ch))
		} else if IsLetter(string(ch)) {
			interval = interval + string(ch)
		} else if IsNumber(string(ch)) || string(ch) == "." {
			interval = interval + string(ch)
		}
	}
	if interval != "" {
		result = append(result, interval)
	}
	return result
}

func IsNumber(n string) bool {
	_, err := strconv.ParseFloat(n, 64)
	if err != nil {
		return false
	}
	return true
}
