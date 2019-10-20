package main

import (
	"fmt"
)

var numbers = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "onethousand",
}

var length = 0

func main() {

	//1 through 20
	for i := 1; i <= 20; i++ {
		addWordLength(numbers[i])
	}

	//21 to 99
	for i := 21; i <= 99; i++ {
		text := spellTens(i)
		addWordLength(text)
	}

	//100
	addWordLength("one" + numbers[100])

	//101-999
	for i := 101; i <= 999; i++ {
		text := spellHundreds(i)
		addWordLength(text)
	}

	//1000
	addWordLength(numbers[1000])

	fmt.Println(length)

}

func spellTens(number int) string {

	ten := ((number / 10) % 10) * 10
	unit := number % 10

	return string(numbers[ten] + numbers[unit])
}

func spellHundreds(number int) string {

	str := ""

	hundred := (number / 100) % 10
	ten := ((number / 10) % 10) * 10
	unit := number % 10

	str += numbers[hundred] + "hundred"

	if ten == 10 {
		str += "and" + numbers[10+unit]

		return str
	}

	if ten == 0 && unit == 0 {
		return str
	}

	if ten == 0 && unit != 0 {
		str += "and" + numbers[unit]

		return str
	}

	if ten != 0 && unit == 0 {
		str += "and" + numbers[ten]

		return str
	}

	if ten != 0 && unit != 0 {
		str += "and" + numbers[ten] + numbers[unit]
		return str
	}

	return str

}

func addWordLength(word string) {

	length += len(word)
}
