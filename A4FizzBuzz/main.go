package main

import (
	"fmt"
)

func main() {

	fmt.Println("Bem vindo ao programa FizzBuzz, por favor informe um valor inicial:")
	var primeiroValor int
	fmt.Scanln(&primeiroValor)
	fmt.Println("Agora informe um valor final:")
	var valorFinal int
	fmt.Scanln(&valorFinal)
	if valorFinal <= primeiroValor {
		fmt.Println("O valor final deve ser maior que o número inicial.")
	}

	fizzBuzz(primeiroValor, valorFinal)
}
