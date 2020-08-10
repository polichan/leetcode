### 1.两数之和

##### 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

#### 示例:
```
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```
方法 1 - 暴力法
```go
func twoSum(nums []int, target int) []int {
	var index1, index2 int
	var ans []int
	for index1 < len(nums) && index2 < len(nums) {
		if index1 < index2 {
			if nums[index1]+nums[index2] == target {
				ans = append(ans, index1)
				ans = append(ans, index2)
				break
			}
		}
		index2 ++
		if index2 == len(nums) {
			index1 ++
			index2 = 0
		}
	}
	return ans
}
```
这个解法没什么可以讲的，让我们看方法 2

方法 2 - HashMap
```go
func twoSumUseMap(nums []int, target int) []int {
	m := map[int]int{}
	for i,v := range nums{
		if k, ok := m[target-v]; ok {
			return []int{i,k}
		}
		m[v] = i
	}
	return nil
}
```
1. 首先我们将 nums 中的每一个元素都以 nums 的 value 为 key，值为 index，即 nums 下标存入。
2. 其次我们在 for 循环体中，每次加入一个新的 nums 中的 value 都去检查 target - 当前循环到的 value，判断是不是存在符合题目的解。
3. 如果存在，则返回。

总结：
1. 暴力法
2. 使用哈希表

延申：
1. 关于 [哈希表](https://github.com/polichan/leetcode/tree/master/src/common/HashMap/README.md "哈希表")