package main

import (
	"fmt"
)

func main() {

	f := largestFactor(600851475143)

	fmt.Println(f)

}

func largestFactor(number int) int {

	var factors []int

	for i := 2; i*i < number; i++ {
		if number%i == 0 {

			if isPrime(i) {
				factors = append(factors, i)
			}

		}
	}

	return factors[len(factors)-1]
}

func isPrime(number int) bool {
	for i := 2; i*i < number; i++ {
		if number%i == 0 {
			return false
		}

	}

	return true

}
