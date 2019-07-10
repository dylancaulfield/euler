package main

import (
	"fmt"
	"go-euler/timing"
)

func main() {

	timing.Start()

	for i := 1; i > 0; i++ {

		tri := nthTrianglular(float64(i))

		if numFactors(int(tri)) >= 500 {

			fmt.Println(int(tri))
			return

		}

	}

	timing.End()

}

func numFactors(num int) int {

	factors := 0

	last := 0

	for f := 1; f*f <= num; f++ {

		if num%f == 0 {

			factors++

			last = f
		}

	}

	factors = factors * 2

	if last*last == num {
		factors--
	}

	return factors

}

func nthTrianglular(n float64) float64 {

	return (0.5 * (n * n)) + (0.5 * n)

}

//2a = 1
//a = 0.5
//
//3(0.5) + b = 3 - 1
//1.5 + b = 2
// b = 0.5
//
//0.5 + 0.5 + c = 1
//c = 0
