package main

func main() {
	ans := minAddToMakeValid("((")
	print(ans)
}

func minAddToMakeValid(S string) int {
	ans := 0
	s := make([]string, 0)
	b := []rune(S)
	for i := 0; i < len(b); i++ {
		currentStr := string(b[i])
		if currentStr == "(" {
			s = append(s, "(")
		}else if len(s) == 0{
			ans++
		}else{
			// pop
			s = append(s[:len(s) - 1],s[len(s):]...)
		}
	}
	return ans + len(s)
}