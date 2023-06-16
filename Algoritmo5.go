package main

import (
	"fmt"
)

func ObterFrequenciaCaracteres(s string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, char := range s {
		frequencia[char]++
	}

	return frequencia
}

func main() {
	s := "Hello, World!"

	frequencia := ObterFrequenciaCaracteres(s)

	for char, freq := range frequencia {
		fmt.Printf("Caractere: %c, Frequência: %d\n", char, freq)
	}
}
