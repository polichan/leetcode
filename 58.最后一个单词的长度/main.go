package main

func main() {
	println(lengthOfLastWord("a "))
}


func lengthOfLastWord(s string) int {
	r := []rune(s)
	if r == nil || len(r) == 0 {
		return 0
	}
	count := 0
	for i := len(r) - 1; i >= 0; i -- {
		cursor := string(r[i])
		if cursor == " "{
			if count == 0 {
				continue
			}
			break
		}
		count ++
	}
	return count
}
