package main

import "fmt"

func main() {

	sundaysCount := 0
	date := Date{1, 1, 1, 1900}

	for {

		if date.dayWeek == 7 && date.dayMonth == 1 && date.year >= 1901 {
			sundaysCount++
		}

		if (date.dayMonth == 31) && (date.month == 12) && (date.year == 2000) {
			break
		}

		date.next()

	}

	fmt.Println(sundaysCount)

}

type Date struct {
	dayWeek  int
	dayMonth int
	month    int
	year     int
}

func (d *Date) next() {

	//Day of Week
	if d.dayWeek == 7 {
		d.dayWeek = 1
	} else {
		d.dayWeek++
	}

	switch d.month {

	//February
	case 2:

		if d.dayMonth < 28 {
			d.dayMonth++
		} else if d.dayMonth == 28 {

			if d.year%4 == 0 {
				d.dayMonth++
			} else {
				d.dayMonth = 1
				d.month++
			}

		} else if d.dayMonth == 29 {
			d.dayMonth = 1
			d.month++
		}

		return

	//April June September November
	case 4, 6, 9, 11:

		if d.dayMonth == 30 {
			d.dayMonth = 1
			d.month++
		} else {
			d.dayMonth++
		}

		return

	//January March May July October
	case 1, 3, 5, 7, 8, 10:

		if d.dayMonth == 31 {
			d.dayMonth = 1
			d.month++
		} else {
			d.dayMonth++
		}

		return

	//December
	case 12:

		if d.dayMonth == 31 {
			d.dayMonth = 1
			d.month = 1
			d.year++
		} else {
			d.dayMonth++
		}

		return
	}

}
