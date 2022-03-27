# go_iter
Go iter tools (for iterating , mapping, filtering, reducing streams -represented as channels-)

- to install : go get -u github.com/serge-hulne/go_iter 
- **Documentation : See in the Documentation directory.**

Defines:

- Filter
- Map
- Reduce
- Range
- Take.

Can easily be extended/generalized to all collection types.

See **examples** for more information:

Partial example (code snippet) :

In the example, data coming from an input channel are mapped/filtered, using Map(), to an output channel.
Map.is one of the functions provided by go_iter.

```	
// Callbacks:

	map_to_square := func(c1, c2 chan int) (error, chan int) {
		for item := range c1 {
			c2 <- item * item
		}
		return nil, c2
	}

	// Data
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Data -> iterator:
	fmt.Println("- - -")
	generated := Iterable_from_array(nums)

	// Chaining iterators:
	even := Map(generated, func(c1, c2 chan int) (error, chan int) {
		for item := range c1 {
			if item%2 == 0 {
				c2 <- item
			}
		}
		return nil, c2
	})
	even_and_squared := Map(even, map_to_square)

	// Displaying results:
	for item := range even_and_squared {
		fmt.Printf("Gen: %#v\n", item)
	}
```

Alternatively, the callbacks can be defined directly in the Map or Filter iterator (JavaScript style):

```
	even := Filter(generated, func(c1, c2 chan int) chan int {
		for item := range c1 {
			if item%2 == 0 {
				c2 <- item
			}
		}
		return c2
	})
```

