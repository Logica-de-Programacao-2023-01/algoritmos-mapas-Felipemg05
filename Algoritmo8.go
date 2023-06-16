package main

import (
	"fmt"
)

func CalcularContaCompartilhada(gastos map[string]float64) map[string]float64 {
	totalGastos := 0.0

	for _, valor := range gastos {
		totalGastos += valor
	}

	mediaGastos := totalGastos / float64(len(gastos))

	resultado := make(map[string]float64)
	for nome, valor := range gastos {
		resultado[nome] = valor - mediaGastos
	}

	return resultado
}

func main() {
	gastos := map[string]float64{
		"João":   100.0,
		"Maria":  150.0,
		"Pedro":  120.0,
		"Carlos": 80.0,
	}

	resultado := CalcularContaCompartilhada(gastos)

	for nome, valor := range resultado {
		fmt.Printf("%s: R$ %.2f\n", nome, valor)
	}
}
