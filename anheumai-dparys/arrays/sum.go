package array

// Sum exposes Array sum= functionality
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll returns for each given slice a checksum
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails returns for each given slice tail a checksum
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

// SumWithoutHeadAndTail checksums each slice without head and tail
func SumWithoutHeadAndTail(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) < 2 {
			sums = append(sums, 0)
		} else {
			tailIndex := len(numbers) - 1
			middle := numbers[1:tailIndex]
			sums = append(sums, Sum(middle))
		}
	}
	return sums
}
