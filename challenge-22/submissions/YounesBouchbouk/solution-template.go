package main

import (
	"fmt"
	"slices"
)

func main() {
	// Standard U.S. coin denominations in cents
	denominations := []int{1, 5, 10, 25, 50}

	// Test amounts
	amounts := []int{87, 42, 99, 33, 7}

	for _, amount := range amounts {
		// Find minimum number of coins
		minCoins := MinCoins(amount, denominations)

		// Find coin combination
		coinCombo := CoinCombination(amount, denominations)

		// Print results
		fmt.Printf("Amount: %d cents\n", amount)
		fmt.Printf("Minimum coins needed: %d\n", minCoins)
		fmt.Printf("Coin combination: %v\n", coinCombo)
		fmt.Println("---------------------------")
	}
}

// MinCoins returns the minimum number of coins needed to make the given amount.
// If the amount cannot be made with the given denominations, return -1.
func MinCoins(amount int, denominations []int) int {

	if amount == 0 || len(denominations) == 0 {
		return 0
	}
	slices.Sort(denominations)

	count := 0
	index := len(denominations) - 1

	for amount >= 0 && index >= 0 {

		if amount-denominations[index] >= 0 {
			amount = amount - denominations[index]
			count++
		} else {
			index--
		}

		if amount == 0 {
			return count
		}
	}
	// TODO: Implement this function
	return -1
}

// CoinCombination returns a map with the specific combination of coins that gives
// the minimum number. The keys are coin denominations and values are the number of
// coins used for each denomination.
// If the amount cannot be made with the given denominations, return an empty map.
func CoinCombination(amount int, denominations []int) map[int]int {

	result := make(map[int]int, 0)

	if amount == 0 || len(denominations) == 0 {
		return result
	}
	slices.Sort(denominations)

	count := 0
	index := len(denominations) - 1

	for amount >= 0 && index >= 0 {

		if amount-denominations[index] >= 0 {
			amount = amount - denominations[index]
			count++
			result[denominations[index]]++
		} else {
			index--
		}

		if amount == 0 {
			return result
		}
	}
	// TODO: Implement this function
	return result
}
