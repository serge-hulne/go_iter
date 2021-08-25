package main

import (
	"fmt"
	. "iter_int"
)

func Fib(n int) Chan {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}

func main_int() {

	// // Enumerate

	for item := range Enumerate(Fib(10)) {
		fmt.Printf("enumerate : %v, %3.2d\n", item.Index, item.Value)
	}

	println("\n - - - ")

	//range
	for item := range Range(10) {
		fmt.Printf("range : %d\n", item)
	}

	println("\n - - - ")

	// //take
	for item := range Take(Fib(10), 3) {
		fmt.Printf("take : %3.2d\n", item)
	}

	//reduce
	println("\n - - - ")

	r := Reduce(Fib(10), func(in Chan) int {
		s := 0
		for x := range in {
			s += x
		}
		return s
	})
	fmt.Printf("reduced : %3.2d\n", r)

	println("\n - - - ")

	mp := Map(Fib(10), func(in Chan, out Chan) Chan {
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

	for i := 0; i < N_channel; i++ {
		channel_array[i] = Map(Fib(10_000_000_000), func(in Chan, out Chan) Chan {
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
