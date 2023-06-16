package main

import (
	"fmt"
	"strings"
)

func ContarLetrasPorPalavra(frase string) map[string]map[rune]int {
	resultado := make(map[string]map[rune]int)

	palavras := strings.Fields(frase)

	for _, palavra := range palavras {
		contagem := make(map[rune]int)

		for _, char := range palavra {
			contagem[char]++
		}

		resultado[palavra] = contagem
	}

	return resultado
}

func main() {
	frase := "A linguagem Go é incrível!"

	resultado := ContarLetrasPorPalavra(frase)

	for palavra, contagem := range resultado {
		fmt.Printf("Palavra: %s\n", palavra)
		for letra, qtd := range contagem {
			fmt.Printf("  Letra: %c, Quantidade: %d\n", letra, qtd)
		}
		fmt.Println()
	}
}
