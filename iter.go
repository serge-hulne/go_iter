package main

/*
	TODO:

		- From_file().
		- From_url()
		- GroupBy()
*/

// Models a pair {index, Value}
type Pair[T comparable] struct {
	Index int
	Value T
}

// Creates an Iterable (channel) from a Slice / Array of data of type [T]
func Iterable_from_Array[T comparable](array []T) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, x := range array {
			out <- x
		}
	}()
	return out
}

//
type Generator[U any] interface {
	Next()
	HasNext() bool
	Value() U
}

// Returns the values from a generator via a channel
func Generator_to_Iterator[U any](c Generator[U]) chan U {
	ch := make(chan U)
	go func() {
		defer close(ch)
		for c.HasNext() {
			ch <- c.Value()
			c.Next()
		}
	}()
	return ch
}

// Maps input channel in to output channel out using callback
func Map[T any](ch1 chan T, f func(T) T) chan T {
	ch2 := make(chan T)
	go func() {
		defer close(ch2)
		for x := range ch1 {
			ch2 <- f(x)
		}
	}()
	return ch2
}

// Filters input channel in to output channel out using callback
func Filter[T any](ch1 chan T, f func(T) bool) chan T {
	ch2 := make(chan T)
	go func() {
		defer close(ch2)
		for x := range ch1 {
			if f(x) {
				ch2 <- x
			}
		}
	}()
	return ch2
}

// Reduces a list of values to a single value (functional)
func Reduce[T any](ch1 chan T, f func(T, T) T) T {
	var temp T
	for x := range ch1 {
		temp = f(temp, x)
	}
	return temp
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

// Skips next N items in a list
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

// Takes the 'nmax' first entries form 'in'
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
