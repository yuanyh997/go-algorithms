package main

import (
	"fmt"
)

//升序排列
func insert_sort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			fmt.Println(arr)
		}

	}
	return arr
}

//插入排序
func insert_sort2(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}

		}
		fmt.Println(arr)

	}
	return arr
}

//冒泡排序
func bubble_sort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//鸡尾酒排序 反复冒泡
func cocktail_sort(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--
		for j := right; j > left; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		left++
	}
	return arr
}

//希尔排序
func shell_sort(arr []int) []int {
	n := len(arr)
	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && arr[j] < arr[j-key] {
				arr[j], arr[j-key] = arr[j-key], arr[j]
				j = j - key
			}
		}
		key = key / 2
	}
	return arr
}

//选择排序
func select_sort(arr []int) []int {
	length := len(arr)
	min := 0
	for i := 0; i < length; i++ {
		min = i
		for j := i; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

//快速排序
func quick_sort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid := arr[0]
	head, tail := 0, len(arr)-1
	for i := 1; i <= tail; {
		if arr[i] > mid {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}
	quick_sort(arr[:head])
	quick_sort(arr[head+1:])
	//return arr
}

func partition(arr []int, low, high int) int {
	mid := arr[low]
	head, tail := low, high
	for {
		for arr[tail] > mid {
			tail--
		}

		for tail > head && arr[head] <= mid {
			head++
		}
		if head >= tail {
			break
		}
		arr[head], arr[tail] = arr[tail], arr[head]

	}
	arr[low], arr[head] = arr[head], arr[low]
	return head
}

//二路快排
func quick_sort_twoway(arr []int, low, high int) {
	if low < high {
		m := partition(arr, low, high)
		quick_sort_twoway(arr, low, m-1)
		quick_sort_twoway(arr, m+1, high)
	}

}

func partition2(arr []int, low, high int) int {
	mid := arr[low]
	head, tail := low, high
	q, p := low, high
	for {
		for arr[tail] >= mid {
			if arr[tail] == mid {
				arr[tail], arr[p] = arr[p], arr[tail]
				p--
			}
			tail--
		}

		for tail > head && arr[head] <= mid {
			if arr[head] == mid {
				arr[head], arr[q] = arr[q], arr[head]
				q++
			}
			head++
		}
		if head >= tail {
			break
		}
		arr[head], arr[tail] = arr[tail], arr[head]

	}
	arr[low], arr[head] = arr[head], arr[low]
	return head
}

//迭代版本 归并排序
func merge_sort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	key := n / 2
	left := merge_sort(arr[:key])
	right := merge_sort(arr[key:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp

}

//递归版本 归并排序
func merge_sort2(arr []int) []int {
	sum := len(arr)
	if sum <= 1 {
		return arr
	}
	left := arr[0 : sum/2]
	lSize := len(left)
	if lSize >= 2 {
		left = merge_sort2(left)
	}
	right := arr[sum/2:]
	rSize := len(right)
	if rSize >= 2 {
		right = merge_sort2(right)
	}
	i := 0
	j := 0
	arr = make([]int, sum)
	for k := 0; k < sum; k++ {
		if i < lSize && j < rSize {
			if left[i] < right[j] {
				arr[k] = left[i]
				i++
			} else {
				arr[k] = right[j]
				j++
			}
		} else if i >= lSize {
			arr[k] = right[j]
			j++
		} else if j >= rSize {
			arr[k] = left[i]
			i++
		}
	}
	return arr
}

//计数排序 多个相同元素
func count_sort(arr []int) []int {
	m := make(map[int]int)
	tmp := make([]int, len(arr))
	max, min := arr[0], arr[0]
	for _, i := range arr {
		m[i] += 1
		if max < i {
			max = i
		}
		if min > i {
			min = i
		}
	}
	fmt.Println(min, max)
	j := 0
	for i := min; i <= max; i++ {
		for m[i] > 0 {
			tmp[j] = i
			j++
			m[i]--
		}
	}
	return tmp
}

func main() {
	//arr := []int{2, 3, 1, 6, 5, 9, 8, 7, 4}
	//ans := bubble_sort(arr)
	//ans := insert_sort2(arr)
	//ans := cocktail_sort(arr)
	//ans := shell_sort(arr)
	//quick_sort(arr)
	//quick_sort_twoway(arr, 0, len(arr)-1)
	//ans := merge_sort(arr)
	//ans := merge_sort2(arr)
	arr2 := []int{2, 2, 1, 3, 1, 4, 2, 7, 4}
	ans := count_sort(arr2)
	fmt.Println(ans)
}
