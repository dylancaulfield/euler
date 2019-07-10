package main

import "fmt"

func main() {

	for c := 3; c > 0; c++ {

		for b := 2; b < c; b++ {

			for a := 1; a < b; a++ {

				if (a*a+b*b == c*c) && a+b+c == 1000 {

					fmt.Println(a * b * c)

					return
				}

			}

		}

	}

}
