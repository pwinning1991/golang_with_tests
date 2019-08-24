package main

func Sum(nums []int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum
}
