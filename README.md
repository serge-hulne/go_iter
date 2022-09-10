# go_iter
Go iter tools (for iterating , mapping, filtering, reducing streams -represented as channels-)

*Thank you to Agustín Díaz for his contributions.*

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
	// Data
	nums := []int{0, 1, 2, 3, 4, 5}

	// Data -> iterator:
	generated := Iterable_from_Array(nums)

	// -> 0, 1, 2, 3, 4, 5

	// Mapping x -> 2 * x for all elements of the iterator:
	even := Map(generated, func(x int) int {
		return 2 * x
	})

	// -> 0, 2, 4, 6, 8, 10

	// Mapping all the elements of 'even' to it's square value:
	even_and_squared := Map(even, func(x int) int {
		return x * x
	})

	// -> 0, 4, 16, 36, 64, 100

	println("Final")

	// Displaying results:
	for item := range even_and_squared {
		fmt.Printf("%#v ", item)
	}
```

