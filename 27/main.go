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
	primes = map[int]bool{}
)

func isPrime(number int) bool {

	if number <= 1 {
		return false
	}

	if primes[number] {
		return true
	}

	for i := 2; i*i < number; i++ {
		if number%i == 0 {

			primes[number] = false
			return false

		}
	}

	primes[number] = true
	return true

}
