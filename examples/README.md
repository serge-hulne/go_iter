```

VARIABLES

var Filter = Map
    Filter : synonymous of Map() -in this model-.


TYPES

type Chan chan Generic
    models a Stream of data (as a channel of Generics)
    
 type Generic interface{}
    Alias for inerface{}

type Pair struct {
	Index int
	Value Generic
}
    Models a pair {index, Value}

type ReduceCallback func(Chan) Generic
    Type of callack func required by Reduce()
    
type Generator interface {
	Next() Generic
	HasNext() bool
}    
    
    
FUNCTIONS

func Enumerate(in Chan) chan Pair
    Lists the elements from 'in' ino 'out' with an index (as a 'Pair')

func Range(nmax int) chan int
    Iterable channel of integers < nmax: Emulates 'range(nmax)' from Python

func Reduce(in Chan, cb ReduceCallback) interface{}
    Reduce (as in other functional programming schemes)

func Every(in Chan, n int) Chan
    Every : Take every in N item from input channel (backpressure management)
    Ex: Every(in, 2) takes every second item from 'in, put sit into into 'out'

func Iterable_from_array(array []Generic) Chan
    Creates an Iterable (channel) from a slice of data of tye [] interface{}

func Iterable_from_generator(gen Generator) Chan
    Creates a Iterable (channel) from generator interface

func Map(in Chan, cb FilterCallback) Chan
    Maps input channel in to output channel out using callback 'cb' of type:
    'ReduceCallback'

func Skip(in Chan, n int) Chan
    Every : Skips N item from input channel (backpressure management) Ex:
    Every(in, 2) skips 2 items after every item read from 'in'

func Slice(in Chan, nmin, nmax int) Chan
    Takes a slice [nmin, nmax] from 'in' into 'out'

func Take(in Chan, nmax int) Chan
    Takes the 'nmax' fist entries form 'in'

type FilterCallback func(Chan, Chan) Chan
    Type of callack func required by Filter() and Map()


```
    
