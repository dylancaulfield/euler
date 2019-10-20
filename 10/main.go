package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}

func isPrime(number int) bool {
	for i := 2; i*i < number; i++ {
		if number%i == 0 {
			return false
		}

	}

	return true

}
