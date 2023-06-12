package main

import "fmt"

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
