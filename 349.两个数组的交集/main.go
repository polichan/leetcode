package main

func main() {
	num1 := []int{1,1,2,2}
	num2 := []int{2,2}
	intersection(num1, num2)
}

func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	if len(nums1) == 0 || len(nums2) == 0 {
		return ans
	}
	hash := make(map[int]bool)
	for _, e := range nums1{
		hash[e] = true
	}
	for _, e := range nums2{
		if hash[e] {
			ans = append(ans, e)
		}
	}
	ans = removeDups(ans)
	return ans
}

func removeDups(elements []int)(nodups []int)  {
	encountered := make(map[int]bool)
	println(encountered)
	for _, element := range elements{
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return nodups
}