package arraysandslices

func Sum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		if len(slice) <= 1 {
			sums = append(sums, 0)
			continue
		}

		slice = slice[1:]

		sums = append(sums, Sum(slice))
	}

	return sums
}
