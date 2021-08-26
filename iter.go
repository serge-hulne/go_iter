package go_iter

// Alias for inerface{}
type Generic interface{}

// Models a pair {index, Value}
type Pair struct {
	Index int
	Value Generic
}

// models a Stream of data (as a channel of Generics)
type Chan chan Generic

// Type of callack func required by Filter() and Map()
type FilterCallback func(Chan, Chan) Chan

// Type of callack func required by Reduce()
type ReduceCallback func(Chan) Generic

// Maps input channel in to output channel out using
// callback 'cb' of type: 'ReduceCallback'
func Map(in Chan, cb FilterCallback) Chan {
	out := make(Chan)
	go func() {
		defer close(out)
		out = cb(in, out)
	}()
	return out
}

// Creates a channel from a slice of data of tye [] interface{}
func Iterable_from_array(array []Generic) Chan {
	out := make(Chan)
	go func() {
		defer close(out)
		for _, x := range array {
			out <- x
		}
	}()
	return out
}

/// Filter : synonymous of Map() -in this model-.
var Filter = Map

// Every : Take every in N item from input channel (backpressure management)
// Ex: Every(in, 2) takes every second item from 'in, put sit into into 'out'
func Every(in Chan, n int) Chan {
	out := make(Chan)
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
func Skip(in Chan, n int) Chan {
	out := make(Chan)
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
func Reduce(in Chan, cb ReduceCallback) interface{} {
	word := cb(in)
	return word
}

// Takes the 'nmax' fist entries form 'in'
func Take(in Chan, nmax int) Chan {
	out := make(Chan)
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
func Slice(in Chan, nmin, nmax int) Chan {
	out := make(Chan)
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
func Enumerate(in Chan) chan Pair {
	out := make(chan Pair)
	go func() {
		defer close(out)
		index := 0
		for i := range in {
			out <- Pair{index, i}
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
