package go_iter

import (
	"fmt"
	. "github.com/serge-hulne/go_iter"
)


// Implements the interface generator
type Counter struct {
	value int
}

func (c *Counter) Next() {
	if c.HasNext() {
		c.value += 1
	}
}

func (c *Counter) HasNext() bool {
	if c.Value() < 10 {
		return true
	} else {
		return false
	}
}

func (c *Counter) Value() int {
	return c.value
}

func main() {

	c := Counter{0}

	it := Generator_to_Iterator[int](&c)

	for v := range it {
		fmt.Printf("%d ", v)
	}

	// - - -

	it = Generator_to_Iterator[int](&c)

	it1 := Map[int](it, func(x int) int {
		return 2 * x
	})

	for v := range it1 {
		fmt.Printf("%d ", v)
	}

	println()

	// - -  -

	println("\n- - -")

	it = Generator_to_Iterator[int](&c)

	it1 = Filter(it, func(x int) bool {
		return x%2 == 0
	})
	
	defer Clear(it1)
	
	for v := range it1 {
		fmt.Printf("%d ", v)
	}

	// - - -

	println("\n- - -")

	it = Generator_to_Iterator[int](&c)

	it1 = Filter(it, func(x int) bool {
		return x%2 == 0
	})

	sum := Reduce(it1, func(x int, y int) int {
		return x + y
	})

	fmt.Printf("Sum = %d", sum)

	println("\n- - -")

	it = Generator_to_Iterator[int](&c)

	every := Every(it, 3)

	for v := range every {
		fmt.Printf("%d ", v)
	}

	println("\n- - -")
	it = Generator_to_Iterator[int](&c)

	skipped := Skip(it, 4)

	for v := range skipped {
		fmt.Printf("%d ", v)
	}

	println("\n- - -")

	it = Generator_to_Iterator[int](&c)

	top := Take(it, 5)

	for v := range top {
		fmt.Printf("%d ", v)
	}

	// - - -

	println("\n- - -")

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

}

