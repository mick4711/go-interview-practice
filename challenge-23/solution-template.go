package main

import (
	"fmt"
	"index/suffixarray"
	"strings"
)

func main() {
	// Sample texts and patterns
	testCases := []struct {
		text    string
		pattern string
	}{
		{"ABABDABACDABABCABAB", "ABABCABAB"},
		{"AABAACAADAABAABA", "AABA"},
		{"GEEKSFORGEEKS", "GEEK"},
		{"AAAAAA", "AA"},
	}

	// Test each pattern matching algorithm
	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Text: %s\n", tc.text)
		fmt.Printf("Pattern: %s\n", tc.pattern)

		// Test naive pattern matching
		naiveResults := NaivePatternMatch(tc.text, tc.pattern)
		fmt.Printf("Naive Pattern Match: %v\n", naiveResults)

		// Test KMP algorithm
		kmpResults := KMPSearch(tc.text, tc.pattern)
		fmt.Printf("KMP Search: %v\n", kmpResults)

		// Test Rabin-Karp algorithm
		rkResults := RabinKarpSearch(tc.text, tc.pattern)
		fmt.Printf("Rabin-Karp Search: %v\n", rkResults)

		fmt.Println("------------------------------")
	}
}

// NaivePatternMatch performs a brute force search for pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func NaivePatternMatch(text, pattern string) []int {
	// check empty strings
	if text == "" || pattern == "" {
		return []int{}
	}

	// get lens and make result slice
	lt := len(text)
	lp := len(pattern)
	res := make([]int, 0, lt/lp)

	// loop thru text char by char
	for i := 0; i <= lt-lp; i++ {
		// check each char in pattern
		k := i
		for j := 0; j < lp; {
			if text[k:k+1] == pattern[j:j+1] {
				// chars match move along
				k++
				j++
			} else {
				// chars don't match stop this pattern match
				break
			}
		}

		// check if pattern matched by count of loop
		if k-i == lp {
			res = append(res, i)
		}
	}
/*
	// loop thru text - std library
	for i := 0; i <= lt-lp; {
		// check for pattern
		j := strings.Index(text[i:], pattern)
		if j >= 0 {
			res = append(res, j+i)
			i += j + 1
		} else {
			i++
		}
	}

*/
	return res
}

// KMPSearch implements the Knuth-Morris-Pratt algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func KMPSearch(text, pattern string) []int {
	// TODO: Implement this function
	// check empty strings
	if text == "" || pattern == "" {
		return []int{}
	}

	// get lens and make result slice
	lt := len(text)
	lp := len(pattern)
	res := make([]int, 0, lt/lp)

	// loop thru text
	for i := 0; i <= lt-lp; {
		// check for pattern
		j := strings.Index(text[i:], pattern)
		if j >= 0 {
			res = append(res, j+i)
			i += j + 1
		} else {
			i++
		}
	}

	return res
}

// RabinKarpSearch implements the Rabin-Karp algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func RabinKarpSearch(text, pattern string) []int {
	// TODO: Implement this function
	// check empty strings
	if text == "" || pattern == "" {
		return []int{}
	}

	// std library suffixarray
	index := suffixarray.New([]byte(text))
	res := index.Lookup([]byte(pattern), -1)
	if res == nil {
		res = []int{}
	}

	return res
}

/*
goos: linux
goarch: amd64
pkg: challenge23
cpu: Intel(R) Core(TM) i7-10610U CPU @ 1.80GHz
- manual char by char
BenchmarkPatternMatching/NaivePatternMatch-8      485155     2910 ns/op     352 B/op    1 allocs/op
- strings.index
BenchmarkPatternMatching/NaivePatternMatch-8     2837085      428.7 ns/op   352 B/op    1 allocs/op
- index/suffixarray
BenchmarkPatternMatching/NaivePatternMatch-8       90572    11592 ns/op    2400 B/op    4 allocs/op

BenchmarkPatternMatching/KMPSearch-8              395563     2837 ns/op     352 B/op    1 allocs/op
BenchmarkPatternMatching/RabinKarpSearch-8        413664     2813 ns/op     352 B/op    1 allocs/op

*/
