package main

import (
	"fmt"
	. "iter_string"
	"strings"
)

func words() Chan {
	words_array := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "fence"}
	words_chan := make(Chan)
	go func() {
		defer close(words_chan)
		for _, w := range words_array {
			words_chan <- w
		}
	}()
	return words_chan
}

func main_string() {

	words_chan := words()

	// Take
	for item := range Take(words_chan, 2) {
		fmt.Printf("take : %v\n", item)
	}

	println(" - - -")

	words_chan = words()

	// Enumerate
	for item := range Enumerate(words_chan) {
		fmt.Printf("enumerate : %d, %v\n", item.Index, item.Value)
	}

	println(" - - -")

	// Range
	for n := range Range(10) {
		fmt.Printf("%v \n", n)
	}

	println(" - - -")

	words_chan = words()

	// Filter
	for item := range Filter(words_chan, filter_cb) {
		fmt.Printf("filtered: %v \n", item)
	}

	println(" - - - ")

	words_chan = words()

	// Map (Map and Filter use Map())
	for item := range Map(words_chan, map_cb) {
		fmt.Printf("Mapped: %v \n", item)
	}

	println(" - - - ")

	words_chan = words()

	// Map (Map and Filter use Map())
	reduced := Reduce(words_chan, reduce_cb)
	fmt.Printf("Reduced: %v \n", reduced)

	println(" - - - ")

}

// Callback for Reduce():
func reduce_cb(in Chan) string {
	word := ""
	for item := range in {
		word = word + "#" + item
	}
	return word
}

// Callback for Filter():
func filter_cb(in Chan, out Chan) Chan {
	for word := range in {
		if len(word) > 3 {
			out <- word
		}
	}
	return out
}

// Callback for Map():
func map_cb(in Chan, out Chan) Chan {
	for word := range in {
		out <- strings.ToUpper(word)
	}
	return out
}
