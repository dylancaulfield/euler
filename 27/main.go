package main

import "fmt"

func main() {

	maxLength := 0
	var maxQuad Quadratic

	for a := -999; a < 1000; a++ {

		for b := -1000; b <= 1000; b++ {

			for n := 0; true; n++ {

				quad := Quadratic{n, a, b}
				value := quad.getValue()

				if !isPrime(value) {
					break
				}

				if n > maxLength {
					maxLength = n
					maxQuad = quad
				}

			}

		}

	}

	fmt.Println(maxQuad.a * maxQuad.b)

}

type Quadratic struct {
	n, a, b int
}

func (q *Quadratic) getValue() int {
	return q.n*q.n + q.a*q.n + q.b
}

var (
	// 0 is unchecked, 1 is checked but not prime, 2 is checked and prime
	primes = map[int]int{}
)

func isPrime(number int) bool {

	if number <= 1 {
		return false
	}

	if primes[number] == 1 {
		return false
	}

	if primes[number] == 2 {
		return true
	}

	for i := 2; i*i < number; i++ {
		if number%i == 0 {

			primes[number] = 1
			return false

		}
	}

	primes[number] = 2
	return true

}
