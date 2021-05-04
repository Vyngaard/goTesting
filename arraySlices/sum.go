package main

func Sum(numbers []int) int {

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var tailSums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		tailSums = append(tailSums, Sum(tail))
	}

	return tailSums
}
