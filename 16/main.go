package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	bigInt := big.NewInt(2)

	bigInt.Exp(bigInt, big.NewInt(1000), big.NewInt(0))

	s := []byte(bigInt.String())

	sum := 0

	for i := 0; i < len(s); i++ {

		integer, _ := strconv.Atoi(string(s[i]))

		sum += integer
	}

	fmt.Println(sum)
}
