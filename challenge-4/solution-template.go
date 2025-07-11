package main

import "fmt"

// ConcurrentBFSQueries concurrently processes BFS queries on the provided graph.
// - graph: adjacency list, e.g., graph[u] = []int{v1, v2, ...}
// - queries: a list of starting nodes for BFS.
// - numWorkers: how many goroutines can process BFS queries simultaneously.
//
// Return a map from the query (starting node) to the BFS order as a slice of nodes.
// YOU MUST use concurrency (goroutines + channels) to pass the performance tests.
func ConcurrentBFSQueries(graph map[int][]int, queries []int, numWorkers int) map[int][]int {
	// TODO: Implement concurrency-based BFS for multiple queries.
	// Return an empty map so the code compiles but fails tests if unchanged.
	return map[int][]int{}
}

// get the path of all nodes connected to the root
func BFS(graph map[int][]int, root int) []int {
	// start with root iem at current level
	Q := []int{root}
	// track all items visited to avoid recursion
	visited := make(map[int]bool)
	visited[root] = true
	// initialise path (output result)
	var path []int

	// iterate thru current level items
	for len(Q) > 0 {
		// take leftmost current item off slice
		current_item := Q[0]
		Q = Q[1:]

		// add top of the Q to the path
		path = append(path, current_item)

		// loop thru adjacent items of current item
		for _, v := range graph[current_item] {
			// if it's an unseen item add to list of items to add to path
			if !visited[v] {
				visited[v] = true
				Q = append(Q, v)
			}
		}
	}

	return path
	/*
			graph
			0: {1, 2},
			1: {2, 3},
			2: {3},
			3: {4},
			4: {},

			get graph index root slice values
		   Possible output:
		   results[0] = [0 1 2 3 4]
		   results[1] = [1 2 3 4]
		   results[2] = [2 3 4]
	*/
}

func main() {
	// You can insert optional local tests here if desired.
	graph := map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {3},
		3: {4},
		4: {},
	}
	// queries := []int{0, 1, 2}
	query := 0

	res := BFS(graph, query)

	fmt.Println(res)
}
