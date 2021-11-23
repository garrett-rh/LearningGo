package v6_main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// New requirement to sum all numbers in the slice besides the first

func SumAllTails(numbersToSum ...[]int) []int { //https://gobyexample.com/variadic-functions: ... allows for any number of trailing arguments
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
