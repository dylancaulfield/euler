package main

import (
	"fmt"
	"math/big"
)

func main() {

	sum := big.NewInt(0)

	for i := int64(1); i <= 1000; i++ {

		toBeAdded := big.NewInt(1)

		j := big.NewInt(i)

		toBeAdded = toBeAdded.Exp(j, j, nil)

		sum = sum.Add(sum, toBeAdded)

	}

	text, _ := sum.MarshalText()

	fmt.Printf("%s", text[len(text) - 10 : ])

}
