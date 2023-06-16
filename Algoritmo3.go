package main

import "fmt"

func SomaValoresMapa(valores map[string]int) int {
	soma := 0

	for _, valor := range valores {
		soma += valor
	}

	return soma
}

func main() {
	valores := map[string]int{
		"valor1": 10,
		"valor2": 20,
		"valor3": 30,
	}

	soma := SomaValoresMapa(valores)

	fmt.Println(soma)
}
