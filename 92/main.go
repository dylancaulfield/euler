package main

import (
	"fmt"
)

func main() {

	arriveTo89 := 0

	for i := 1; i < 10000000; i++ {

		var chain []int
		chain = append(chain, i)

		for {

			digits := getDigits(chain[len(chain)-1])
			digitsSum := 0

			for _, d := range digits {
				digitsSum += d * d
			}

			chain = append(chain, digitsSum)

			// If last in chain is 1,
			if chain[len(chain)-1] == 1 {

				break
			}

			if chain[len(chain)-1] == 89 {

				arriveTo89++
				break
			}

		}

	}

	fmt.Println(arriveTo89)

}

func getDigits(number int) []int {

	var digits []int

	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}

	for i := 0; i < len(digits)/2; i++ {

		digits[i], digits[len(digits)-1-i] = digits[len(digits)-1-i], digits[i]

	}

	return digits

}
