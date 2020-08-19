package main

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(1)
	s.Top()
	s.Pop()
	s.Empty()
}

type Queue struct {
	 data []int
}


func (q *Queue)Push(x int){
	q.data = append(q.data, x)
}

func (q * Queue) Pop() int {
	return q.data[len(q.data) - 1]
}

func (q *Queue)isEmpty()  bool {
	return len(q.data) == 0
}

func (q * Queue) Peek()(exist bool, x int)  {
	if len(q.data) != 0 {
		x = q.data[0]
		q.data = q.data[1:]
		return true, x
	}else {
		return false, 0
	}
}

type MyStack struct {
	q1 Queue
	q2 Queue
	top int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{q1: Queue{data: []int{}}, q2: Queue{data: []int{}}, top: 0}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.top = x
	this.q1.Push(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for len(this.q1.data) > 1 {
		_, pop := this.q1.Peek()
		this.top = pop
		this.q2.Push(pop)
	}

	_, pop := this.q1.Peek()
	temp := this.q1
	this.q1 = this.q2
	this.q2 = temp
	return pop
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.top
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q1.isEmpty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */