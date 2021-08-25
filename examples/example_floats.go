package main

import (
	"fmt"
	. "iter_float"
)

func floats() Chan {
	floats_array := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	floats_chan := make(Chan)
	go func() {
		defer close(floats_chan)
		for _, w := range floats_array {
			floats_chan <- w
		}
	}()
	return floats_chan
}

func main_float() {

	floats_chan := floats()

	// Enumerate
	for item := range Enumerate(floats_chan) {
		fmt.Printf("enumerate : %v, %3.2f\n", item.Index, item.Value)
	}

	println("\n - - - ")
	//range
	for item := range Range(10) {
		fmt.Printf("range : %d\n", item)
	}

	println("\n - - - ")
	floats_chan = floats()

	//take
	for item := range Take(floats_chan, 3) {
		fmt.Printf("take : %3.2f\n", item)
	}

	//reduce

	println("\n - - - ")
	floats_chan = floats()

	r := Reduce(floats_chan, func(in Chan) float64 {
		s := 0.0
		for x := range in {
			s += x
		}
		return s
	})
	fmt.Printf("reduced : %3.2f\n", r)

	println("\n - - - ")
	floats_chan = floats()

	mp := Map(floats_chan, func(in Chan, out Chan) Chan {
		for x := range in {
			out <- x * x
		}
		return out
	})

	for item := range mp {
		fmt.Printf("mapped : %v  \n", item)
	}

	// Parallel Map:
	println("\n - - - ")

	const N_channel = 10
	var channel_array [N_channel]Chan
	floats_chan = floats()
	for i := 0; i < N_channel; i++ {
		channel_array[i] = Map(floats_chan, func(in Chan, out Chan) Chan {
			for x := range in {
				out <- x * x
			}
			return out
		})
	}

	for channel_id, mp := range channel_array {
		for item := range mp {
			fmt.Printf("Parallel cannel_id : %2d, val : %6v \n", channel_id, item)
		}
	}

}
