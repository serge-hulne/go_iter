package go_iter

type Generic interface{}

type Pair struct {
	Index int
	Value Generic
}

type Chan chan Generic

type FilterCallback func(Chan, Chan) Chan

type ReduceCallback func(Chan) Generic

//
func Map(in Chan, cb FilterCallback) Chan {
	out := make(Chan)
	go func() {
		defer close(out)
		out = cb(in, out)
	}()
	return out
}

//
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

//
var Filter = Map

//
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

//
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

//
func Reduce(in Chan, cb ReduceCallback) interface{} {
	word := cb(in)
	return word
}

//
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

//
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

//
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

//
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
