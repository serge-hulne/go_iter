# go_iter
Go iter tools (for iterating , mapping, filtering, reducing streams -represented as channels-)

- to install : go get -u github.com/serge-hulne/go_iter 

Defines:

- Filter
- Map
- Reduce
- Range
- Take.

Can easily be extended/generalized to all collection types.

See examples for more information:

Partial example (code snippet) :

In the example, data coming from an input channel are mapped/filtered, using Map(), to an output channel.
Map.is one of the functions provided by go_iter.

```
	//... User defined type
	type Person struct {
		Name string
		Age  int
	}
	
	//... Send list of "Persons in a channel"
	input_channel := input()

	for item := range input_channel {
		fmt.Printf("%v, %v\n", item.(Person).Name, item.(Person).Age)
	}

	fmt.Println("- - -")
	
	//... Refresh channel
	input_channel = input()

	cb := func(c1, c2 Chan) Chan {
		for person := range c1 {
			//... Some fitering action here:
			c2 <- filtered_item
		}
		return c2
	}

	electors := Map(input_channel, cb)

	for person := range electors {
		fmt.Printf("Elector --> %v\n", person.(Person))
	}
```


go doc 

```
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
    Creates a channel from a slice of data of tye [] interface{}

func Map(in Chan, cb FilterCallback) Chan
    Maps input channel in to output channel out

func Skip(in Chan, n int) Chan
    Every : Skips N item from input channel (backpressure management) Ex:
    Every(in, 2) skips 2 items after every item read from 'in'

func Slice(in Chan, nmin, nmax int) Chan
    Takes a slice [nmin, nmax] from 'in' into 'out'

func Take(in Chan, nmax int) Chan
    Takes the 'nmax' fist entries form 'in'


    
TYPES

type Chan chan Generic
    models a Stream of data (as a channel of Generics

type Generic interface{}
    Alias for inerface{}

type Pair struct {
	Index int
	Value Generic
}
    Models a pair {index, Value}

type ReduceCallback func(Chan) Generic
    Type of callack func required by Reduce()
    
type FilterCallback func(Chan, Chan) Chan
    Type of callack func required by Filter() and Map()

```
