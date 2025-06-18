package xiter_test

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/moeryomenko/xiter"
)

func ExampleIterFunc() {
	// Create a sequence from a slice of integers
	seq := slices.Values([]int{1, 2, 3})

	// Apply a function that squares each element
	squared := xiter.IterFunc(seq, func(e int) int { return e * e })

	fmt.Println(slices.Collect(squared))
	// Output: [1 4 9]
}

func ExampleFilter() {
	// Create a sequence from a slice of integers
	s := []int{1, 2, 3, 4, 5, 6}

	// Filter even numbers from a slice
	filtered := xiter.Filter(s, func(n int) bool { return n%2 == 0 })

	fmt.Println(filtered)
	// Output: [2 4 6]
}

func ExampleFilterSeq() {
	// Create a sequence from a slice of integers
	words := slices.Values([]string{"apple", "banana", "cherry", "date"})

	// Filter strings longer than 5 characters
	filteredWords := xiter.FilterSeq(words, func(s string) bool { return len(s) > 5 })

	fmt.Println(slices.Collect(filteredWords))
	// Output: [banana cherry]
}

func ExampleAppendFunc() {
	// Create a sequence from a slice of integers
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	// Append the elements of s2 to s1 after multiplying by 2
	appended := xiter.AppendFunc(s1, s2, func(v int) int { return v * 2 })

	fmt.Println(appended)
	// Output: [1 2 3 8 10 12]
}

func ExampleAppendIf() {
	// Create a sequence from a slice of integers
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	// append only even numbers from s2 to s1
	appended := xiter.AppendIf(s1, s2, func(v int) bool { return v%2 == 0 })

	fmt.Println(appended)
	// Output: [1 2 3 4 6]
}

func ExampleAppendSeqIf() {
	// Create a sequence from a slice of integers
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	// Append elements from s2 to s1 that are greater than 4
	appended := xiter.AppendSeqIf(s1, slices.Values(s2), func(v int) bool { return v > 4 })

	fmt.Println(appended)
	// Output: [1 2 3 5 6]
}

func ExapleAppendSeqFunc() {
	// Create a sequence from a slice of integers
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	// Append the elements of s2 to s1 after multiplying by 3
	appended := xiter.AppendSeqFunc(s1, slices.Values(s2), func(v int) int { return v * 3 })

	fmt.Println(appended)
	// Output: [1 2 3 12 15 18]
}

func ExampleMap() {
	intSlice := []int{1, 2, 3}

	// Apply a function that convert to string each element
	stringSlice := xiter.Map(intSlice, func(e int) string { return strconv.Itoa(e) })

	fmt.Println(stringSlice)
	// Output: [1 2 3]
}
