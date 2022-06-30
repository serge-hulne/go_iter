```
package main

import (
	"fmt"
	. "github.com/serge-hulne/go_iter"
)

type Counter struct {
	Value   int
	HasNext bool
}

func (c *Counter) Next() {
	if c.Value < 10 {
		c.Value++
		c.HasNext = true
	} else {
		c.HasNext = false
	}
}

func main() {

	c := Counter{0, true}

	it := Generator_to_Iterator(c)

	for v := range it {
		fmt.Printf("%d ", v)
	}

	// - - -

	it = Generator_to_Iterator(c)

	it1 := Map(it, func(x int) int {
		return 2 * x
	})

	for v := range it1 {
		fmt.Printf("%d ", v)
	}

	println()

	// - -  -

	println("\n- - -")

	it = Generator_to_Iterator(c)

	it1 = Filter(it, func(x int) bool {
		return x%2 == 0
	})

	for v := range it1 {
		fmt.Printf("%d ", v)
	}

	// - - -

	println("\n- - -")

	it = Generator_to_Iterator(c)
	it1 = Filter(it, func(x int) bool {
		return x%2 == 0
	})

	sum := Reduce(it1, func(x int, y int) int {
		return x + y
	})

	fmt.Printf("Sum = %d", sum)

	println("\n- - -")

	it = Generator_to_Iterator(c)
	every := Every(it, 3)

	for v := range every {
		fmt.Printf("%d ", v)
	}

	println("\n- - -")

	it = Generator_to_Iterator(c)

	skipped := Skip(it, 4)

	for v := range skipped {
		fmt.Printf("%d ", v)
	}

	println("\n- - -")

	it = Generator_to_Iterator(c)

	top := Take(it, 5)

	for v := range top {
		fmt.Printf("%d ", v)
	}

}

```
