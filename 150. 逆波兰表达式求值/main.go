package main

import "strconv"

func main() {

}

func evalRPN(tokens []string) int {
	number := []int{}
	for _, val := range tokens{
		l := len(number)
		switch val {
		case "+":
			number  = append(number[:l -2], number[l-2] + number[l-1])
		case "-":
			number  = append(number[:l -2], number[l-2] - number[l-1])
		case "*":
			number  = append(number[:l -2], number[l-2] * number[l-1])
		case "/":
			number  = append(number[:l -2], number[l-2] / number[l-1])
		default:
			num, _ := strconv.Atoi(val)
			number  = append(number, num)
		}
	}
	return number[0]
}