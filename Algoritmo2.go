package main

import "fmt"

func MesclarMapas(mapa1, mapa2 map[string]string) map[string]string {
	resultado := make(map[string]string)

	for chave, valor := range mapa1 {
		resultado[chave] = valor
	}

	for chave, valor := range mapa2 {
		resultado[chave] = valor
	}

	return resultado
}

func main() {
	mapa1 := map[string]string{
		"chave1": "valor1",
		"chave2": "valor2",
	}

	mapa2 := map[string]string{
		"chave2": "novoValor2",
		"chave3": "valor3",
	}

	resultado := MesclarMapas(mapa1, mapa2)

	fmt.Println(resultado)
}
