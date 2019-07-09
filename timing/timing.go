package timing

import (
	"fmt"
	"time"
)

var beganSeconds int
var beganNanoSeconds int

func Start() {

	t := time.Now()

	beganSeconds = t.Second()
	beganNanoSeconds = t.Nanosecond()

}

func End() {

	t := time.Now()

	endSeconds := t.Second()
	endNanoSeconds := t.Nanosecond()

	takenSeconds := endSeconds - beganSeconds
	takenNanoSeconds := endNanoSeconds - beganNanoSeconds

	fmt.Println(
		"Time taken:", takenSeconds, "seconds or",
		takenNanoSeconds/1000000, "milliseconds or",
		takenNanoSeconds, "nanoseconds")

}
