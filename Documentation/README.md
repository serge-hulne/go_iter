```
FUNCTIONS

func Enumerate[T comparable](in chan T) chan Pair[T]
    Lists the elements from 'in' ino 'out' with an index (as a 'Pair')

func Every[T comparable](in chan T, n int) chan T
    Every : Take every in N item from input channel (backpressure management)
    Ex: Every(in, 2) takes every second item from 'in, put sit into into 'out'

func Filter[T any](ch1 chan T, f func(T) bool) chan T
    Filters input channel in to output channel out using callback

func Generator_to_Iterator[U any](c Generator[U]) chan U
    Returns the values from a generator via a channel

func Iterable_from_Array[T comparable](array []T) chan T
    Creates an Iterable (channel) from a Slice / Array of data of type [T]

func Map[T any](ch1 chan T, f func(T) T) chan T
    Maps input channel in to output channel out using callback

func Range(nmax int) chan int
    Iterable channel of integers < nmax: Emulates 'range(nmax)' from Python

func Reduce[T any](ch1 chan T, f func(T, T) T) T
    Reduces a list of values to a single value (functional)

func Skip[T comparable](in chan T, n int) chan T
    Skips next N items in a list

func Slice[T comparable](in chan T, nmin, nmax int) chan T
    Takes a slice [nmin, nmax] from 'in' into 'out'

func Take[T comparable](in chan T, nmax int) chan T
    Takes the 'nmax' first entries form 'in'


TYPES

type Counter struct {
	// Has unexported fields.
}
    Implements the interface generator

func (c *Counter) HasNext() bool

func (c *Counter) Next()

func (c *Counter) Value() int

type Generator[U any] interface {
	Next()
	HasNext() bool
	Value() U
}

type Pair[T comparable] struct {
	Index int
	Value T
}
    Models a pair {index, Value}

```
