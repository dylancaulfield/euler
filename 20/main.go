package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	prod := big.NewInt(1)

	for n := int64(100); n > 0; n-- {

		prod.Mul(prod, big.NewInt(n))

	}

	bytes := []byte(prod.String())

	var ints = make([]int, len(bytes))

	for i, s := range bytes {

		integer, err := strconv.Atoi(string(s))
		if err != nil {
			panic(err)
		}

		ints[i] = integer
	}

	sum := 0

	for i := 0; i < len(ints); i++ {
		sum += ints[i]
	}

	fmt.Println(sum)

}
