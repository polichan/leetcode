package main

func main() {
	subsets([]int{
		1, 2, 3,
	})
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var backtrack func(path []int, start int)
	backtrack = func(path []int, start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(path, i+1)
			path = path[:len(path)-1]
		}

	}
	backtrack([]int{}, 0)
	return res
}
