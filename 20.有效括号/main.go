package main

import "fmt"

func main() {
	isValid("(())")
}

func isValid(s string) bool {
	r := []rune(s)
	stack := CreateStack(len(r))
	for i := 0; i < len(r); i++ {
		currentChar := string(r[i])
		if currentChar == "(" {
			stack.Push(")")
		}else if currentChar == "[" {
			stack.Push("]")
		}else if currentChar == "{" {
			stack.Push("}")
		}else if stack.IsEmpty() || currentChar != stack.Pop() {
			return false
		}
	}
	return stack.IsEmpty()
}

type Stack struct {
	size int
	top int
	data []interface{}
}

func CreateStack(size int) Stack{
	s := Stack{}
	s.size = size
	s.top = -1
	s.data = make([]interface{}, size)
	return s
}
func (s *Stack) IsEmpty() bool  {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.size -1
}

func (s *Stack)Push(data interface{}) bool {
	if s.IsFull() {
		panic("栈已经满")
	}
	s.top ++
	// 把当前元素放在栈顶
	s.data[s.top] = data
	return true
}

func (s *Stack)GetLength() int  {
	length := s.top + 1
	return length
}

func (s *Stack) Clear()  {
	s.top = -1
}


// pop,返回栈顶元素
func (s *Stack) Pop() interface{}  {
	// 判断是否是空栈
	if s.IsEmpty() {
		fmt.Println("stack is empty , pop error")
		return nil
	}
	// 把栈顶的元素赋值给临时变量tmp
	tmp := s.data[s.top]
	// 栈顶指针-1
	s.top--
	return tmp
}