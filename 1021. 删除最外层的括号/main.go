package main

import (
	"fmt"
)

func main(){
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}

func removeOuterParentheses(S string) string {
	level, res := 0, ""
	r := []rune(S)
	for i := 0; i < len(r); i++{
		t := string(r[i])
		if t == "(" {
			level++
			if level == 1 {
				continue
			}
		}else{
			level --
			if level == 0 {
				continue
			}
		}
		res += t
	}
	return res
}

