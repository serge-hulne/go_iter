package main

import (
	"fmt"
	. "go_iter"

	//. "github.com/serge-hulne/go_iter"

	"strings"
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

	cb := func(c1, c2 chan Person) chan Person {
		for person := range c1 {
			if person.Age > 18 {
				p := Person{
					strings.ToUpper(person.Name),
					person.Age}
				c2 <- p
			}
		}
		return c2
	}

	electors := Filter(input_channel, cb)

	for person := range electors {
		fmt.Printf("Elector --> %v\n", person)
	}

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("- - -")
	gen := Iterable_from_array(nums)

	for item := range gen {
		fmt.Printf("Gen: %#v\n", item)
	}

	fmt.Println("- - -")
	gen = Iterable_from_array(nums)

	every := Every(gen, 2)
	for item := range every {
		fmt.Printf("Every: %#v\n", item)
	}

	fmt.Println("- - -")
	gen = Iterable_from_array(nums)

	skipped := Skip(gen, 2)
	for item := range skipped {
		fmt.Printf("Skipped: %#v\n", item)
	}

	fmt.Println("- - -")
	gen = Iterable_from_array(nums)

	slice := Slice(gen, 4, 5)
	for item := range slice {
		fmt.Printf("Slice: %#v\n", item)
	}

	fmt.Println("- - -")

}
