package main

import (
	"fmt"
	"go-euler/timing"
)

func main() {

	timing.Start()

	sum := 1
	previous := 1

	var fibEven []int

	for sum <= 4000000 {

		x := sum

		sum = sum + previous

		previous = x

		if sum%2 == 0 {
			fibEven = append(fibEven, sum)
		}
	}

	evenSum := 0

	for _, num := range fibEven {
		evenSum += num
	}

	fmt.Println(evenSum)

	timing.End()

}
