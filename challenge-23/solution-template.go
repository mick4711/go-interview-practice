package main

import (
	"fmt"
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

	// loop thru text
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

	return res
}

// KMPSearch implements the Knuth-Morris-Pratt algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func KMPSearch(text, pattern string) []int {
	// TODO: Implement this function
	return nil
}

// RabinKarpSearch implements the Rabin-Karp algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func RabinKarpSearch(text, pattern string) []int {
	// TODO: Implement this function
	return nil
}
