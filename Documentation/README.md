
```
FUNCTIONS

func Enumerate[T comparable](in chan T) chan Pair[T]
    Lists the elements from 'in' ino 'out' with an index (as a 'Pair')

func Every[T comparable](in chan T, n int) chan T
    Every : Take every in N item from input channel (backpressure management)
    Ex: Every(in, 2) takes every second item from 'in, put sit into into 'out'

func Filter[T comparable](in chan T, cb FilterCallback[T]) chan T
func Iterable_from_array[T comparable](array []T) chan T
    Creates an Iterable (channel) from a slice of data of type [T]

func Map[T comparable](in chan T, cb FilterCallback[T]) chan T
    Maps input channel in to output channel out using callback 'cb' of type:
    'ReduceCallback'

func Range(nmax int) chan int
    Iterable channel of integers < nmax: Emulates 'range(nmax)' from Python

func Reduce[T comparable](in chan T, cb ReduceCallback[T]) T
    Reduce (as in other functional programming schemes)

func Skip[T comparable](in chan T, n int) chan T
    Every : Skips N item from input channel (backpressure management) Ex:
    Every(in, 2) skips 2 items after every item read from 'in'

func Slice[T comparable](in chan T, nmin, nmax int) chan T
    Takes a slice [nmin, nmax] from 'in' into 'out'

func Take[T comparable](in chan T, nmax int) chan T
    Takes the 'nmax' fist entries form 'in'


TYPES

type FilterCallback[T comparable] func(chan T, chan T) chan T
    Type of callack func required by Filter() and Map()

type Pair[T comparable] struct {
	Index int
	Value T
}
    Models a pair {index, Value}

type ReduceCallback[T comparable] func(chan T) T
    Type of callack func required by Reduce()
    
```
