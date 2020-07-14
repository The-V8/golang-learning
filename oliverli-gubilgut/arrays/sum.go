package main

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
