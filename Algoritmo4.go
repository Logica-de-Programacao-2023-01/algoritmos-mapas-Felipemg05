package main

import (
	"fmt"
	"sort"
	"strings"
)

func ObterAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {
		chave := ordenarString(palavra)
		anagramas[chave] = append(anagramas[chave], palavra)
	}

	return anagramas
}

func ordenarString(s string) string {
	letras := strings.Split(s, "")
	sort.Strings(letras)
	return strings.Join(letras, "")
}

func main() {
	palavras := []string{"amor", "roma", "carro", "arco", "mar", "rota", "macaco"}

	gruposAnagramas := ObterAnagramas(palavras)

	for _, grupo := range gruposAnagramas {
		fmt.Println(grupo)
	}
}
