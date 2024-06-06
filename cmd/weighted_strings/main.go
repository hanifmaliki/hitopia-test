package main

import (
	"fmt"
)

func calculateWeights(s string) map[int]bool {
	// I use map[int]bool instead []int because i don't need to loop Weights to check Query is included or not
	weights := map[int]bool{}
	n := len(s)
	i := 0

	for i < n {
		charWeight := int(s[i] - 'a' + 1)
		currentWeight := charWeight
		weights[currentWeight] = true

		// Get weight char and substring
		for j := i + 1; j < n && s[j] == s[i]; j++ {
			currentWeight += charWeight
			weights[currentWeight] = true
		}

		// Shift char until the char is different
		i++
		for i < n && s[i] == s[i-1] {
			i++
		}
	}

	return weights
}

func checkQueries(weights map[int]bool, queries []int) []string {
	results := make([]string, len(queries))

	for i, query := range queries {
		if weights[query] {
			results[i] = "Yes"
		} else {
			results[i] = "No"
		}
	}

	return results
}

func weightedStrings(s string, queries []int) []string {
	weights := calculateWeights(s)
	return checkQueries(weights, queries)
}

func main() {
	s := "abbcccdee" // 1, 2, 4, 3, 6, 9, 4, 5, 10
	queries := []int{1, 3, 9, 8, 16, 2, 8, 5}
	results := weightedStrings(s, queries)
	fmt.Println(results)
}
