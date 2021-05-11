package main

import (
	"fmt"
	"strings"
)

func main() {
	target := []int{
		2,3,4,
	}
	fmt.Println(strings.Join(buildArray(target, 4), ","))
}

func buildArray(target []int, n int) []string {

	length := len(target)
	index := 0

	operations := []string{}
	for i:=1; i<=n; i++ {
		if index >= length {
			break
		}

		if target[index] == i {
			operations = append(operations, "Push")
			index++
			continue
		}
		operations = append(operations, "Push", "Pop")
	}
	return operations
}