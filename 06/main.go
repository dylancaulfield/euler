package main

import (
	"fmt"
)

func main() {

	//Sum of squares
	sum := sumOfSquares(1, 100)

	// Square of sums
	square := squareOfSums(1, 100)

	fmt.Println(square - sum)

}

func sumOfSquares(lower, upper int) int {

	sum := 0

	for i := lower; i <= upper; i++ {

		sum += i * i

	}

	return sum

}

func squareOfSums(lower, upper int) int {

	sum := 0

	for i := lower; i <= upper; i++ {

		sum += i

	}

	sum = sum * sum

	return sum

}
