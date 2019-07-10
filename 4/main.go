package main

import (
	"fmt"
	"go-euler/timing"

	"strconv"
)

func main() {

	timing.Start()

	largest := 1000

	for i := 100; i < 1000; i++ {

		for j := 100; j < 1000; j++ {

			x := i * j

			isPal := isPalindrome(x)

			if x > largest && isPal {
				largest = x
			}

		}

	}

	fmt.Println(largest)

	timing.End()

}

func isPalindrome(num int) bool {

	str := strconv.Itoa(num)

	rev := make([]byte, len(str))

	for i := len(str) - 1; i >= 0; i-- {

		rev[len(str)-i-1] = str[i]
	}

	if string(rev) == str {
		return true
	}

	return false

}
