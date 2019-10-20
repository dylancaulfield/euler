package main

import (
	"fmt"
	"strconv"
)

func main() {

	circularsCount := 0

	for i := 2; i < 100; i++ {

		if !isPrime(i) {
			continue
		}

		var arr [][]int
		arr = append(arr, getDigits(i))

		permutations := getPermutations(arr)

		permutationsWithoutDuplicates := removeDuplicates(permutations)

		for _, p := range permutationsWithoutDuplicates {

			if !isPrime(formNumber(p)) {
				continue
			}

			// All are prime..
			circularsCount++
		}

	}

	fmt.Println(circularsCount)

}

// TODO: check if already exists
func getPermutations(numbers [][]int) [][]int {

	if len(numbers[0]) == 2 {

		var integers [][]int

		integers = append(integers, []int{numbers[0][0], numbers[0][1]})
		integers = append(integers, []int{numbers[0][1], numbers[0][0]})

		return integers
	}

	var argument [][]int

	//each number

	for i := 0; i < len(numbers); i++ {

		// each digit of number

		for j := 0; j < len(numbers[i]); j++ {

			swapped := swap(numbers[i], j)

			// get slice from 2nd element on and recurse

			var numbersFrom2ndIndex [][]int

			numbersFrom2ndIndex = append(numbersFrom2ndIndex, swapped[1:])

			permutations := getPermutations(numbersFrom2ndIndex)

			// add 1st piece back on

			for _, perm := range permutations {

				perm = append([]int{swapped[0]}, perm...)

				argument = append(argument, perm)

			}

		}

	}

	// TODO: remove duplicates

	return argument

}

func getDigits(number int) []int {

	var numbers []int

	s := strconv.Itoa(number)

	for _, n := range s {

		integer, _ := strconv.Atoi(string(n))

		numbers = append(numbers, integer)

	}

	return numbers

}

func formNumber(nums []int) int {

	number := 0

	i := 1

	for j := len(nums) - 1; j >= 0; j-- {

		number += nums[j] * i

		i *= 10

	}

	return number

}

func swap(numbers []int, index int) []int {

	var newSlice []int
	newSlice = append(newSlice, numbers[index])

	for i := 0; i < len(numbers); i++ {

		if i == index {
			continue
		}

		newSlice = append(newSlice, numbers[i])

	}

	return newSlice

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

func removeDuplicates(numbers [][]int) [][]int {

	store := map[int]bool{}
	var nonDuplicated [][]int

	for _, n := range numbers{

		number := formNumber(n)

		if store[number] {
			continue
		}

		store[number] = true

		digits := getDigits(number)
		nonDuplicated = append(nonDuplicated, digits)

	}

	return nonDuplicated
}
