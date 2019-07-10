package main

import (
	"fmt"
)

func main() {

	p := nthPrime(6)

	fmt.Println(p)

}

func nthPrime(n int) int {

	primes := 0

	for i := 2; primes <= n; i++ {

		if isPrime(i) {
			primes++

			fmt.Println(i, primes)
		}

		if primes == n {
			return i
		}

	}

	return 0
}

func isPrime(number int) bool {

	for i := 2; i*i < number; i++ {
		if number%i == 0 {
			return false
		}

	}

	return true

}
