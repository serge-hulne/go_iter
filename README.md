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
