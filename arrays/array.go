package arrays

func ArraySum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func ArraySumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = ArraySum(numbers)
	}

	return sums
}
