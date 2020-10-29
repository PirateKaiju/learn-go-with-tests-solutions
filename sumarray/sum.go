package sumarray

//Sum returns the sum of an array
func Sum(numbers [3]int) int {
	sum := 0
	/*for i := 0; i < 3; i++ {
		sum += numbers[i]
	}*/

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlice(numbers)
	}

	return sums

}
