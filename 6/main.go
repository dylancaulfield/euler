package main

import (
	"fmt"
	"github.com/dyldawg/euler/timing"
)

func main() {

	timing.Start()

	//Sum of squares
	sum := sumOfSquares(1, 100)

	// Square of sums
	square := squareOfSums(1, 100)

	fmt.Println(square - sum)


	timing.End()

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
