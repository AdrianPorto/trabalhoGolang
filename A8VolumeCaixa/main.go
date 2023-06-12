package main

import (
	"fmt"
)

// Crie um programa que calcule e apresente o valor do volume de uma caixa retangular.
// VOLUME = COMPRIMENTO * LARGURA * ALTURA.

func main() {
	var comprimento, largura, altura, volume float64

	println("-----------------------------------")
	println("| Informe o comprimento do caixa: |")
	println("-----------------------------------")
	fmt.Scanln(&comprimento)

	println("-----------------------------------")
	println("| Informe o largura do caixa:     |")
	println("-----------------------------------")
	fmt.Scanln(&largura)

	println("-----------------------------------")
	println("| Informe o altura do caixa:      |")
	println("-----------------------------------")
	fmt.Scanln(&altura)

	volume += comprimento * largura * altura

	println("-----------------------------------")
	fmt.Printf("Volume da caixa: %.2f\n", volume)
	println("-----------------------------------")

}
