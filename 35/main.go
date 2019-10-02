package main

import (
	"strconv"
)

func main() {

	

}

func permutations(numbers [][]int) [][]int {

	if len(numbers[0]) == 2 {

		var integers [][]int

		integers = append(integers, []int{numbers[0][0], numbers[0][1]})
		integers = append(integers, []int{numbers[0][1], numbers[0][0]})

		return integers
	}

	var argument [][]int

	//each number

	for i := 0; i < len(numbers); i++ {

		// each number of number

		for j := 0; j < len(numbers[i]); j++ {

			swapped := swap(numbers[i], j)

			// get slice from 2nd element on and recurse

			var numbersFrom2ndIndex [][]int

			numbersFrom2ndIndex = append(numbersFrom2ndIndex, swapped[1:])

			perms := permutations(numbersFrom2ndIndex)

			// add 1st piece back on

			for _, perm := range perms {

				var with1st []int
				with1st = append(with1st, swapped[0])

				for _, ints := range perm {
					with1st = append(with1st, ints)
				}

				argument = append(argument, with1st)

			}

		}

	}

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
