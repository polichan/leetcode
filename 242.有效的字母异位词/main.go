package main

func main() {
	println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	m := make(map[string]int)
	a := []rune(s)
	b := []rune(t)
	if len(b) < len(a) {
		return false
	}
	for _,c := range a{
		_, ok  := m[string(c)]
		if ok {
			m[string(c)] ++
		}else {
			m[string(c)] = 1
		}
	}
	for _, d := range b{
		_, ok  := m[string(d)]
		if !ok {
			return false
		}else{
			if m[string(d)] <= 0 {
				return false
			}else {
				m[string(d)] --
			}
		}

	}
	return true
}