package main

/**
* https://miryang.gitbook.io/learn-go-with-tests/go/arrays-and-slices
**/

// func Sum(numbers []int) int {
// 	sum := 0

// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}

// 	return sum
// }

// refactor
func Sum(numbers []int) int {
	sum := 0
	// _ = idx , numer = value
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

//     for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

// refactor
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

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