package main

import (
	"fmt"
	"math/big"
)

func main() {

	current := big.NewInt(1)
	last := big.NewInt(0)
	next := big.NewInt(0)


	for i := 2; true; i++ {

		next = next.Add(last, current)
		last.Set(current)
		current.Set(next)

		bytes, _ := current.MarshalText()

		length := len([]rune(string(bytes)))

		if length == 1000 {

			fmt.Println(i)

			return
		}

	}

}