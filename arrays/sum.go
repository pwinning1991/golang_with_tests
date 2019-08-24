package main

func Sum(nums []int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
