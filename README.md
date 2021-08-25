# go_iter
Go iter tools (for iterating , mapping, filtering, reducing streams -represented as channels-)

- to install : go get -u github.com/serge-hulne/go_iter 

Defines:

- Filter
- Map
- Reduce
- Range
- Take on collections of integers.

Can easily be extended/generalized to all collection types.

See examples for more information:

Partial example (code snippet) :

```
	input_channel := input()

	for item := range input_channel {
		fmt.Printf("%v, %v\n", item.(Person).Name, item.(Person).Age)
	}

	fmt.Println("- - -")
	input_channel = input()

	cb := func(c1, c2 Chan) Chan {
		for person := range c1 {
			if person.(Person).Age > 18 {
				p := Person{
					strings.ToUpper(person.(Person).Name),
					person.(Person).Age}
				c2 <- p
			}
		}
		return c2
	}

	electors := Map(input_channel, cb)

	for person := range electors {
		fmt.Printf("Elector --> %v\n", person.(Person))
	}
```
