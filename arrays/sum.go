package main

func Sum(nums [5]int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum
}
