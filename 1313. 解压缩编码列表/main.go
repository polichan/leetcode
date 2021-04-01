package main

func main() {
	
}

/**
给你一个以行程长度编码压缩的整数列表 nums 。
考虑每对相邻的两个元素 [freq, val] = [nums[2*i], nums[2*i+1]] （其中 i >= 0 ），每一对都表示解压后子列表中有 freq 个值为 val 的元素，你需要从左到右连接所有子列表以生成解压后的列表。
请你返回解压后的列表。
 */
func decompressRLElist(nums []int) []int {
	left, right := 0, len(nums)
	if right == 0 {
		return []int{}
	}
	var res []int
	for left < right{
		res = append(res, generate(nums[left], nums[left+1])...)
		left = left + 2
	}
	return res
}
// 多少个 Y
func generate(x, y int)[]int  {
	res := make([]int, x)
	for index, _ := range res{
		res[index] = y
	}
	return res
}