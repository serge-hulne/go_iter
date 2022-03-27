package go_iter

import (
	"log"
)

/*
	TODO:

		- From_file().
		- From_url()
		- Make an example for reduce().
		- GroupBy()
*/

// Models a pair {index, Value}
type Pair[T comparable] struct {
	Index int
	Value T
}

// Type of callack func required by Filter() and Map()
type FilterCallback[T comparable] func(chan T, chan T) (error, chan T)

// Type of callack func required by Reduce()
type ReduceCallback[T comparable] func(chan T) (error, T)

// Maps input channel in to output channel out using
// callback 'cb' of type: 'ReduceCallback'
func Map[T comparable](in chan T, cb FilterCallback[T]) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		var e error
		e, out = cb(in, out)
		if e != nil {
			log.Panicf("Encountered error: %s\n", e)
		}
	}()
	return out
}

// func Filter[T comparable](in chan T, cb FilterCallback[T]) chan T {
// 	out := make(chan T)
// 	go func() {
// 		defer close(out)
// 		out = cb(in, out)
// 	}()
// 	return out
// }

// Creates an Iterable (channel) from a slice of data of type [T]
func Iterable_from_array[T comparable](array []T) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, x := range array {
			out <- x
		}
	}()
	return out
}

// Every : Take every in N item from input channel (backpressure management)
// Ex: Every(in, 2) takes every second item from 'in, put sit into into 'out'
func Every[T comparable](in chan T, n int) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		index := 0
		for x := range in {
			if index%n == 0 {
				out <- x
				index = 0
			}
			index++
		}
	}()
	return out
}

// Every : Skips  N item from input channel (backpressure management)
// Ex: Every(in, 2) skips 2 items after every item read from 'in'
func Skip[T comparable](in chan T, n int) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		index := 0
		n++
		for x := range in {
			if index == n || index == 0 {
				out <- x
				index = 0
			}
			index++
		}
	}()
	return out
}

// Reduce (as in other functional programming schemes)
func Reduce[T comparable](in chan T, cb ReduceCallback[T]) T {
	var e error
	e, word := cb(in)
	if e != nil {
		log.Fatalf("Encountered error: %s\n", e)
	}
	return word
}

// Takes the 'nmax' fist entries form 'in'
func Take[T comparable](in chan T, nmax int) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		index := 0
		for i := range in {
			index++
			if index <= nmax {
				out <- i
			} else {
				break
			}
		}
	}()
	return out
}

// Takes a slice [nmin, nmax] from 'in' into 'out'
func Slice[T comparable](in chan T, nmin, nmax int) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		index := 0
		for i := range in {
			if index < nmin {
				index++
				continue
			} else if index <= nmax && index >= nmin {
				out <- i
			} else if index > nmax {
				break
			}
			index++
		}
	}()
	return out
}

//  Lists the elements from 'in' ino 'out' with an index (as a 'Pair')
func Enumerate[T comparable](in chan T) chan Pair[T] {
	out := make(chan Pair[T])
	go func() {
		defer close(out)
		index := 0
		for i := range in {
			out <- Pair[T]{index, i}
			index++
		}
	}()
	return out
}

// Iterable channel of integers < nmax: Emulates 'range(nmax)' from Python
func Range(nmax int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for index := 0; index < nmax; index++ {
			out <- index
		}
	}()
	return out
}
