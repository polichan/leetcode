package main

func main() {
	
}

// num1 1 2 3 0 0 m 5
// num2 2 3 n = 2
func merge(nums1 []int, m int, nums2 []int, n int)  {
	for m > 0 && n > 0 {
		if nums1[m - 1] > nums2[n - 1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
}

func max(x, y int)int  {
	if x > y {
		return x
	}
	return y
}