package main

import "fmt"

func CountPairs(numbers []int) map[[2]int]int {
	pairs := make(map[[2]int]int)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			pairs[pair]++
		}
	}

	return pairs
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 2, 3, 4, 4}

	pairFrequencies := CountPairs(numbers)

	for pair, frequency := range pairFrequencies {
		fmt.Printf("Pair %v - Frequency: %d\n", pair, frequency)
	}
}
