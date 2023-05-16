package main

import "fmt"

func countChars(str string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range str {
		charCount[char]++
	}
	return charCount
}

func main() {
	str := "abracadabra"
	charCount := countChars(str)
	for char, count := range charCount {
		fmt.Printf("'%c' ocorre %d vezes na string.\n", char, count)
	}
}
