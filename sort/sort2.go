package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

func times(arr []int) int {
	max := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	time := 0
	for max > 0 {
		max /= 10
		time++
	}
	return time

}
func radix_sort(arr []int) []int {
	idx := 0
	t := times(arr)
	res := 1
	for l := 0; l < t; l++ {
		link := make([][]int, 10)
		for _, i := range arr {
			idx = (i / res) % 10
			link[idx] = append(link[idx], i)
		}
		k := 0
		for i := 0; i < 10; i++ {
			for _, j := range link[i] {
				arr[k] = j
				k++
			}
		}
		fmt.Println(arr)
		res *= 10
	}
	return arr
}
func main() {

	head := node{2, nil}
	fmt.Println(head)
	arr := []int{2, 3, 1, 6, 5, 9, 8, 7, 4, 4}
	radix_sort(arr)
	arr1 := []int{20001, 2000, 200, 234, 23, 28, 8, 2002, 203}
	fmt.Println(radix_sort(arr1))
	fmt.Println(times(arr1))
}
