package main

import "fmt"

func main() {
	nums := getNums(10)
	for _, num := range nums {
		var word string
		// where is ternary if 0_0
		if num%2 == 0 {
			word = "even"
		} else {
			word = "odd"
		}
		fmt.Println(num, "is", word)
	}
}

func getNums(n int) []int {
	var nums []int
	for i := 0; i <= n; i++ {
		nums = append(nums, i)
	}
	return nums
}
