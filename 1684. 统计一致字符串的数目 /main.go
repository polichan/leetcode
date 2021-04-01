package main

func main() {

}


func countConsistentStrings(allowed string, words []string) int {
	if len(words) == 0 {
		return 0
	}
	m := map[string]int{}
	a := []rune(allowed)
	count := 0
	for _, i := range a {
		m[string(i)]++
	}
	for _, item := range words{
		isCurrentSValid := true
		for _, s := range []rune(item){
			if m[string(s)]  == 0 {
				isCurrentSValid = false
				break
			}
		}
		if isCurrentSValid {
			count++
		}
	}
	return count
}