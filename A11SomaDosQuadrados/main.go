package main

import "fmt"

//Pedra-papel-tesoura-lagarto-Spock é uma adaptação do JoKenPo na série The Big Bang Theory.
//➢ Crie um programa que seja um juiz de Jokenpo (Sheldon Cooper) que dada a jogada dos dois jogadores informa o
//resultado da partida.
//➢ As regras são as seguintes:
//○ Tesoura corta papel
//○ Papel cobre pedra
//○ Pedra esmaga lagarto
//○ Lagarto envenena Spock
//○ Spock esmaga (ou derrete) tesoura
//○ Tesoura decapita lagarto
//○ Lagarto come papel
//○ Papel refuta Spock
//○ Spock vaporiza pedra
//○ Pedra amassa tesoura

func main() {
	var valor1, valor2, valor3 int

	println("----------------------------")
	println("| Digite o primeiro valor: |")
	println("----------------------------")
	fmt.Scanln(&valor1)

	println("----------------------------")
	println("| Digite o segundo valor:  |")
	println("----------------------------")
	fmt.Scanln(&valor2)

	println("----------------------------")
	println("| Digite o terceiro valor: |")
	println("----------------------------")
	fmt.Scanln(&valor3)

	somaDosQuadrados := valor1*valor1 + valor2*valor2 + valor3*valor3

	println("----------------------------------------------------")
	fmt.Printf(" A soma dos quadrados dos três valores é: %d\n", somaDosQuadrados)
	println("----------------------------------------------------")
}
