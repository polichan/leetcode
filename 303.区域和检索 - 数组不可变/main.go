package main

func main() {
	
}

type NumArray struct {
	Value []int
}

func Constructor(nums []int) NumArray {
	arr := NumArray{[]int{0}}  // 浪费第一个空间
	// 1, 2, 3, 4, 5
	for i, v := range nums {
		// 1, 3, 6, 10, 15
		arr.Value = append(arr.Value, v + arr.Value[i])
	}
	return arr
}

func (this *NumArray) SumRange(i int, j int) int {
	// i = 0; j = 3
	// [3+1] - [0] = 15 - 1 = 14 = 1 + 2 + 3 + 4
	return this.Value[j+1] - this.Value[i]
}
