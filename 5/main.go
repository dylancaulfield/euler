package main

import (
	"fmt"
	"go-euler/timing"
)

func main() {

	timing.Start()

	x := smallestMultiple(1, 20)

	fmt.Println(x)

	timing.End()

}

func smallestMultiple(lower, upper int) int {

	for i := 1; i > 0; i++ {

		p := 0

		for j := lower; j <= upper; j++ {

			if i%j == 0 {
				p++
			}

			if p == upper {

				return i
			}

		}

		p = 0

	}

	return 0

}
