package main

import "fmt"

func main() {
	fmt.Print(letterCombinations("23"))
}


func letterCombinations(digits string) []string {
	m := buildNumberMap()
	if digits == "" {
		return []string{}
	}
	result := []string{""}

	for _, val := range digits{
		chars := m[string(val)]
		tempResult := appendChar(result, chars)
		result = tempResult
	}
	return result
}

func appendChar(lastResult []string, currentString []string)[]string  {
	result := make([]string, 0)
	for _, val := range currentString{
		for _, r := range lastResult{
			result = append(result, r + val)
		}
	}
	return result
}

func buildNumberMap()map[string][]string{
	m := make(map[string][]string)
	m["1"] = []string{}
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
	return m
}