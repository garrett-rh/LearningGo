package v4_main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) { //https://gobyexample.com/variadic-functions: ... allows for any number of trailing arguments

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers)) //Takes a slice & returns slice + new items
	}

	return sums
}
