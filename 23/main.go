package main

import "fmt"

// Abuns structure for pointer
type Abuns struct {
	abundants []int
}

func main() {

	var writtenAsSums []int
	abuns := Abuns{getAbundantNums(28123)}

	writtenAsAbundantsMap := writtenAsAbundants(28123, &abuns)

	for i := 0; i < 28123; i++ {

		if !writtenAsAbundantsMap[i] {
			writtenAsSums = append(writtenAsSums, i)
		}

	}

	fmt.Println(getSum(writtenAsSums))

}

func getDivisors(n int) []int {

	var divisors []int

	divisors = append(divisors, 1)

	for i := 2; i*i <= n; i++ {

		if i*i == n {

			divisors = append(divisors, i)

			break
		}

		if n%i == 0 {

			divisors = append(divisors, i, n/i)

		}

	}

	return divisors

}

func getSum(nums []int) int {

	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}

func getAbundantNums(limit int) []int {

	var nums, divisors []int
	var sum int

	for i := 1; i <= limit; i++ {

		divisors = getDivisors(i)
		sum = getSum(divisors)

		if sum > i {
			nums = append(nums, i)
		}

	}

	return nums

}

func writtenAsAbundants(limit int, abundants *Abuns) map[int]bool {

	m := map[int]bool{}

	for i := 0; i < len(abundants.abundants); i++ {

		for j := 0; j < len(abundants.abundants); j++ {

			m[abundants.abundants[i]+abundants.abundants[j]] = true

		}

	}

	return m

}
