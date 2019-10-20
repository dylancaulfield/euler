package main

import (
	"fmt"
)

func main() {

	var div3or5 []int

	for i := 1; i < 1000; i++ {

		if i%3 == 0 || i%5 == 0 {

			div3or5 = append(div3or5, i)

		}

	}

	var sum int

	for _, num := range div3or5 {
		sum += num
	}

	fmt.Println(sum)

}
