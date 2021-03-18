package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf(interpret("G()(al)"))
}

func interpret(command string) string {
	s := make([]string, len(command))
	rS := []rune(command)
	s = append(s, string(rS[0]))
	topItem := string(rS[0])
	for i := 1; i < len(rS); i++{
		item := rS[i]
		sItem := string(item)
		if sItem == ")"{
			if topItem == "(" {
				i := len(s) -1
				// pop 掉， push 进入 o
				s = append(s[:i], s[i+1:]...)
				s = append(s, "o")
				topItem = "o"
				continue
			}else if topItem == "l"{
				i := len(s) - 3
				s = s[:i]
				s = append(s, "al")
				topItem = "l"
				continue
			}
		}
		s = append(s, sItem)
		topItem = sItem
	}
	return strings.Join(s, "")
}
