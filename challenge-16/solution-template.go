package main

import (
	// "bytes"
	"slices"
	"strings"
	"time"
)

// SlowSort sorts a slice of integers using a very inefficient algorithm (bubble sort)
func SlowSort(data []int) []int {
	// Make a copy to avoid modifying the original
	result := make([]int, len(data))
	copy(result, data)

	// Bubble sort implementation
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

// OptimizedSort is your optimized version of SlowSort
// It should produce identical results but perform better
func OptimizedSort(data []int) []int {
	// Hint: Consider using sort package or a more efficient algorithm

	// sort.Slice(data, func(i, j int) bool {
	// 	return data[i] < data[j]
	// })

	slices.Sort(data)

	/*
			   go test -v -timeout 30s -run="SlowSort|OptimizedSort"
			   go test -benchmem -bench="SlowSort|OptimizedSort"

			   BenchmarkSlowSort/10-8           8783551             123.9 ns/op            80 B/op          1 allocs/op
			   BenchmarkSlowSort/100-8           107521             11458 ns/op           896 B/op          1 allocs/op
			   BenchmarkSlowSort/1000-8             607           1653977 ns/op          8192 B/op          1 allocs/op
		sort.Slice
			   BenchmarkOptimizedSort/10-8      9244336             159.2 ns/op            56 B/op          2 allocs/op
			   BenchmarkOptimizedSort/100-8     2408090             532.3 ns/op            56 B/op          2 allocs/op
			   BenchmarkOptimizedSort/1000-8     313731              3719 ns/op            56 B/op          2 allocs/op
		slices.Sort
			   BenchmarkOptimizedSort/10-8     63324622             18.60 ns/op             0 B/op          0 allocs/op
			   BenchmarkOptimizedSort/100-8     9154833             133.4 ns/op             0 B/op          0 allocs/op
			   BenchmarkOptimizedSort/1000-8    1000000              1162 ns/op             0 B/op          0 allocs/op
	*/

	return data
}

// InefficientStringBuilder builds a string by repeatedly concatenating
func InefficientStringBuilder(parts []string, repeatCount int) string {
	result := ""

	for i := 0; i < repeatCount; i++ {
		for _, part := range parts {
			result += part
		}
	}

	return result
}

// OptimizedStringBuilder is your optimized version of InefficientStringBuilder
// It should produce identical results but perform better
func OptimizedStringBuilder(parts []string, repeatCount int) string {
	// Hint: Consider using strings.Builder or bytes.Buffer

	// var b bytes.Buffer // A Buffer needs no initialization.
	var b strings.Builder

	for range repeatCount {
		for _, part := range parts {
			b.Write([]byte(part))
		}
	}

	/*
			go test -v -timeout 30s -run="StringBuilder"
			go test -benchmem -bench="StringBuilder"

			BenchmarkInefficientStringBuilder/Small-8      836750         1385 ns/op        1912 B/op         29 allocs/op
			BenchmarkInefficientStringBuilder/Medium-8       6750       187868 ns/op      518168 B/op        699 allocs/op
			BenchmarkInefficientStringBuilder/Large-8          73     17936059 ns/op    70153736 B/op       7003 allocs/op
		bytes.Buffer
			BenchmarkOptimizedStringBuilder/Small-8       3746812          306.2 ns/op       304 B/op          3 allocs/op
			BenchmarkOptimizedStringBuilder/Medium-8       228163         4949 ns/op        5440 B/op          7 allocs/op
			BenchmarkOptimizedStringBuilder/Large-8         22855        54112 ns/op       84544 B/op         11 allocs/op
		strings.Builder
			BenchmarkOptimizedStringBuilder/Small-8       4266385          289.4 ns/op       248 B/op          5 allocs/op
			BenchmarkOptimizedStringBuilder/Medium-8       344193         3603 ns/op        3320 B/op          9 allocs/op
			BenchmarkOptimizedStringBuilder/Large-8         25576        47915 ns/op       84728 B/op         18 allocs/op
	*/

	return b.String()
}

// ExpensiveCalculation performs a computation with redundant work
// It computes the sum of all fibonacci numbers up to n
// TODO: Optimize this function to be more efficient
func ExpensiveCalculation(n int) int {
	if n <= 0 {
		return 0
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += fibonacci(i)
	}

	return sum
}

// Helper function that computes the fibonacci number at position n
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// OptimizedCalculation is your optimized version of ExpensiveCalculation
// It should produce identical results but perform better
func OptimizedCalculation(n int) int {
	// TODO: Implement a more efficient calculation method
	// Hint: Consider memoization or avoiding redundant calculations
	return ExpensiveCalculation(n) // Replace this with your optimized implementation
}



// TODO try out benchstat tool https://pkg.go.dev/golang.org/x/perf/cmd/benchstat 


// HighAllocationSearch searches for all occurrences of a substring and creates a map with their positions
// TODO: Optimize this function to reduce allocations
func HighAllocationSearch(text, substr string) map[int]string {
	result := make(map[int]string)

	// Convert to lowercase for case-insensitive search
	lowerText := strings.ToLower(text)
	lowerSubstr := strings.ToLower(substr)

	for i := 0; i < len(lowerText); i++ {
		// Check if we can fit the substring starting at position i
		if i+len(lowerSubstr) <= len(lowerText) {
			// Extract the potential match
			potentialMatch := lowerText[i : i+len(lowerSubstr)]

			// Check if it matches
			if potentialMatch == lowerSubstr {
				// Store the original case version
				result[i] = text[i : i+len(substr)]
			}
		}
	}

	return result
}

// OptimizedSearch is your optimized version of HighAllocationSearch
// It should produce identical results but perform better with fewer allocations
func OptimizedSearch(text, substr string) map[int]string {
	// TODO: Implement a more efficient search method with fewer allocations
	// Hint: Consider avoiding temporary string allocations and reusing memory
	return HighAllocationSearch(text, substr) // Replace this with your optimized implementation
}

// A function to simulate CPU-intensive work for benchmarking
// You don't need to optimize this; it's just used for testing
func SimulateCPUWork(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
		// Just waste CPU cycles
		for i := 0; i < 1000000; i++ {
			_ = i
		}
	}
}
