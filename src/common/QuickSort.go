package quicksort

type IQuickSort interface {
	sort()
}

type QuickSort struct {
	arr []int
	start int
	end int
}

func CreateQuickSort(arr []int, start int, end int) QuickSort  {
	qs := QuickSort{arr, start, end}
	return qs
}


/**
快速排序
快速排序通过分支法的思想，从一个数组中选取一个基准元素pivot，
把这个数组中小于pivot的移动到左边，把大于pivot的移动到右边。
然后再分别对左右两边数组进行快速排序。
*/
func (qs QuickSort) Sort(arr []int, start int, end int) {
	if qs.start > qs.end {
		return
	}
	pivotIndex := partition(qs.arr, qs.start, qs.end)
	qs.Sort(qs.arr, qs.start, qs.end - 1)
	qs.Sort(qs.arr, pivotIndex + 1, qs.end)
}

/**
双边循环法
设置两个指针 left 和 right，最初分别指向数组的左右两端。
比较 right 指针指向元素和 pivot 元素，如果 right 元素大于 pivot 元素，right指针左移一位，再和pivot进行比较，如果right元素小于pivot元素的话停止移动，换到left指针。
left指针的操作是，left指针指向的元素和pivot元素比较，如果left指向元素小于或等于pivot，left指针右移，如果left元素大于pivot元素，停止移动。
左右都停止移动后，交换left和right指向的元素，这样left指针指向的是一个小于pivot的元素，right指向的是一个大于pivot的元素。
当left和right重叠的时候结束比较，将第一个元素和left，right指向的元素做交换，完成一轮排序
*/
func partition(arr []int, startIndex int, endIndex int)int  {
	var (
		pivot = arr[startIndex]
		left = startIndex
		right = endIndex
	)

	for left != right {
		for left < right && pivot < arr[right] {
			right --
		}

		for left < right && pivot >= arr[left] {
			left ++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[startIndex], arr[left] = arr[left], arr[startIndex]
	return left
}