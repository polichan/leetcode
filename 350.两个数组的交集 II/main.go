package main

func main() {
	num1 := []int{4,9,5}
	num2 := []int{9,4,9,8,4}
	intersection(num1, num2)
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersection(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1{
		m[num] ++
	}
	var intersection []int
	for _, num := range nums2{
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num] --
		}
	}
	return intersection
}
