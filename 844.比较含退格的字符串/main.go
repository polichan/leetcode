package main

func main() {
	test := "ab#c"
	backspaceCompare(test, test)
}

type Stack struct {
	top string
	data []string
}

func Constructor()Stack  {
	return Stack{top: "", data: []string{}}
}
func (s *Stack)Push(str string)  {
	if s.top == "" && str == "#"{
		return
	}

	if str == "#" {
		s.Pop()
	}else {
		s.data = append(s.data, str)
		s.top = str
	}
}

func (s *Stack)Pop()  {
	l := len(s.data)
	if l >= 1 {
		s.top = s.data[l - 1]
		s.data = append(s.data[:l - 1],s.data[l : ] ...)
	}
}

func (s *Stack)GetAll() (res string) {
	for _,b := range s.data{
		res += b
	}
	return res
}

func backspaceCompare(S string, T string) bool {
	stack := Constructor()
	stack2 := Constructor()
	s := []rune(S)
	t := []rune(T)

	for i := 0; i< len(s); i ++ {
		stack.Push(string(s[i]))
	}

	for i := 0; i< len(t); i ++{
		stack2.Push(string(t[i]))
	}
	return stack.GetAll() == stack2.GetAll()
}