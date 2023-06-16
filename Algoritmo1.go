package main

import (
	"fmt"
	"strings"
)

func ContarPalavras(texto string) map[string]int {
	palavras := strings.Fields(texto)

	ocorrencias := make(map[string]int)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {
	texto := "Olá mundo olá mundo ola mundo"

	ocorrencias := ContarPalavras(texto)

	for palavra, ocorrencia := range ocorrencias {
		fmt.Printf("%s: %d\n", palavra, ocorrencia)
	}
}
