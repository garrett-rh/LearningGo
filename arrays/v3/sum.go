package v3_main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) { //https://gobyexample.com/variadic-functions: ... allows for any number of trailing arguments
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers) //creates a slice of numbers

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
