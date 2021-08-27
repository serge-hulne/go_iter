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

func input() Chan {

	persons := []Generic{
		Person{"Joe", 10},
		Person{"Jane", 40},
		Person{"Jim", 50},
		Person{"John", 60},
	}

	persons_chan := Iterable_from_array(persons)

	return persons_chan
}

// - - -

type Gen struct {
	value Generic
}

func (g *Gen) Next() Generic {
	temp := g.value
	g.value = temp.(int) + 1
	return temp.(int)
}

func (g *Gen) HasNext() bool {
	return g.value.(int) < 10
}

// - - -

func main_generic() {

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

	nums := []Generic{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("- - -")
	gen := Iterable_from_array(nums)

	for item := range gen {
		fmt.Printf("Gen: %#v\n", item.(int))
	}

	fmt.Println("- - -")
	gen = Iterable_from_array(nums)

	every := Every(gen, 2)
	for item := range every {
		fmt.Printf("Every: %#v\n", item.(int))
	}

	fmt.Println("- - -")
	gen = Iterable_from_array(nums)

	skipped := Skip(gen, 2)
	for item := range skipped {
		fmt.Printf("Skipped: %#v\n", item.(int))
	}

	fmt.Println("- - -")
	gen = Iterable_from_array(nums)

	slice := Slice(gen, 4, 5)
	for item := range slice {
		fmt.Printf("Slice: %#v\n", item.(int))
	}

	fmt.Println("- - -")

	// Testing generator struct:
	g := Gen{value: 0}
	generated := Iterable_from_generator(&g)
	for item := range generated {
		fmt.Printf("generated: %v\n", item)
	}

}

/*
	TODO:
		- From_file.
		- From_url
*/
