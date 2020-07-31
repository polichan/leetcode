package main

func main() {
	println(removeDuplicates("abbaca"))
}

type Stack struct {
	top string
	data []string
}

func New()  Stack{
	return Stack{top: "", data: []string{}}
}

func (this *Stack)Push(s string)  {
	if len(this.data) == 0 {
		this.top = s
		this.data = append(this.data, s)
	}else {
		if this.top == s {
			this.Pop()
		}else {
			this.data = append(this.data, s)
			this.top = s
		}
	}
}

func (this *Stack)Pop()string{
	lastIndex := len(this.data) - 1
	if lastIndex - 1 >= 0 {
		this.top = this.data[lastIndex - 1]
	}else {
		this.top = ""
	}
	this.data = append(this.data[:lastIndex],this.data[lastIndex + 1:]...)
	return this.top
}

func (this *Stack)GetAll()string  {
	res := ""
	for i := 0; i < len(this.data); i ++ {
		res += this.data[i]
	}
	return res
}

func removeDuplicates(S string) string {
	str := []rune(S)
	stack := New()
	for i := 0; i< len(str); i ++ {
		stack.Push(string(str[i]))
	}
	return stack.GetAll()
}