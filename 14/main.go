package main

import (
	"fmt"
	"go-euler/timing"
	"math"
	"sync"
)

func main() {

	timing.Start()

	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	maxLength := 0
	maxLengthStarter := float64(0)

	for i := float64(0); i < 1000000; i++ {

		wg.Add(1)
		go chainAndCount(i, &mutex, &wg, &maxLength, &maxLengthStarter)

	}

	wg.Wait()

	fmt.Println(maxLengthStarter)

	timing.End()

}

func even(num float64) float64 {
	return num / 2
}

func odd(num float64) float64 {
	return (3 * num) + 1
}

func chainAndCount(i float64, mutex *sync.Mutex, wg *sync.WaitGroup, maxLength *int, maxLengthStarter *float64) {

	var chain []float64
	// Chaining
	chain = append(chain, i)

	for i > 1 {

		if math.Mod(i, 2) == 0 {
			i = even(i)
			chain = append(chain, i)
		} else {
			i = odd(i)
			chain = append(chain, i)
		}
	}

	mutex.Lock()

	if len(chain) > *maxLength {
		*maxLength = len(chain)
		*maxLengthStarter = chain[0]
	}

	mutex.Unlock()

	wg.Done()

}
