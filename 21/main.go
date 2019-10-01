package main

import "fmt"

func main() {

	amicableSum := 0

	for i := 1; i < 10000; i++ {

		if i == d(d(i)) && i != d(i) {
			amicableSum += i
		}
	}

	fmt.Println(amicableSum)

}

func d(n int) int {

	sum := 0

	for i := 1; i <= n/2; i++ {

		if n%i == 0 {
			sum += i
		}

	}

	return sum

}
