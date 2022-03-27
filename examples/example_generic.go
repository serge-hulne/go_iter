package main

import (
	"fmt"

	"strings"

	. "github.com/serge-hulne/go_iter"
)

type Person struct {
	Name string
	Age  int
}

// - - -

func input() chan Person {

	persons := []Person{
		{"Joe", 10},
		{"Jane", 40},
		{"Jim", 50},
		{"John", 60},
	}

	persons_chan := Iterable_from_array(persons)

	return persons_chan
}

// - - -

func main() {

	input_channel := input()

	for item := range input_channel {
		fmt.Printf("%v, %v\n", item.Name, item.Age)
	}

	fmt.Println("- - -")
	input_channel = input()

	cb := func(c1, c2 chan Person) (error, chan Person) {
		for person := range c1 {
			if person.Age > 18 {
				p := Person{
					strings.ToUpper(person.Name),
					person.Age}
				c2 <- p
			}
		}
		//return errors.New("XXX"), c2
		return nil, c2
	}

	electors := Map(input_channel, cb)

	for person := range electors {
		fmt.Printf("Elector --> %v\n", person)
	}

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

	fmt.Println("- - -")
	generated = Iterable_from_array(nums)

	every := Every(generated, 2)
	for item := range every {
		fmt.Printf("Every: %#v\n", item)
	}

	fmt.Println("- - -")
	generated = Iterable_from_array(nums)

	skipped := Skip(generated, 2)
	for item := range skipped {
		fmt.Printf("Skipped: %#v\n", item)
	}

	fmt.Println("- - -")
	generated = Iterable_from_array(nums)

	slice := Slice(generated, 4, 5)
	for item := range slice {
		fmt.Printf("Slice: %#v\n", item)
	}

	fmt.Println("- - -")

}
