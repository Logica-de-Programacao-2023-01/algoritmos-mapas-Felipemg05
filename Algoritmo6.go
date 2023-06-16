package main

import (
	"fmt"
)

func SomarContagens(textos []map[string]int) map[string]int {
	resultado := make(map[string]int)

	for _, texto := range textos {
		for palavra, contagem := range texto {
			resultado[palavra] += contagem
		}
	}

	return resultado
}

func main() {
	textos := []map[string]int{
		{"hello": 2, "world": 1},
		{"hello": 1, "world": 3},
		{"hello": 4, "world": 2},
	}

	resultado := SomarContagens(textos)

	for palavra, contagem := range resultado {
		fmt.Printf("Palavra: %s, Contagem: %d\n", palavra, contagem)
	}
}
