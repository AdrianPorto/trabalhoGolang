package main

import (
	"fmt"
)

// Crie um programa que receda diversos números que o usuário informar e armazene em um array enquanto o usuário desejar
// inseri-los. Depois, quando o usuário finalizar as inserções, calcule a média dos elementos informados

func main() {
	var numeros []float64

	continuar := true

	for continuar {
		var numero float64

		println("---------------------------------------")
		println("| Digite um número:                   |")
		println("---------------------------------------")
		fmt.Scanln(&numero)

		numeros = append(numeros, numero)

		var opcao string
		println("---------------------------------------")
		println("| Deseja inserir mais números? (S/N): |")
		println("---------------------------------------")
		fmt.Scanln(&opcao)

		if opcao != "S" && opcao != "s" {
			continuar = false
		}
	}

	if len(numeros) > 0 {
		media := calcularMedia(numeros)
		println("---------------------------------------")
		fmt.Printf("A média dos números informados é: %.2f\n", media)
		println("---------------------------------------")
	} else {
		println("---------------------------------------")
		println("| Nenhum número foi informado.        |")
		println("---------------------------------------")
	}
}

func calcularMedia(numeros []float64) float64 {
	total := 0.0

	for _, numero := range numeros {
		total += numero
	}

	media := total / float64(len(numeros))
	return media
}
