package main

import "fmt"

func main() {
	fmt.Println("Código para calcular Soma de Naturais")

	var inicio, fim, soma int

	fmt.Print("Informe o valor do inicio do intervalo: ")
	fmt.Scan(&inicio)

	fmt.Print("\nInforme o valor do final do intervalo: ")
	fmt.Scan(&fim)

	for i := inicio; i < fim; i++ {
		soma += i
	}

	fmt.Printf("O intervalo escolhido foi de %d a %d", inicio, fim)
	fmt.Print("\nSoma: \n", soma)
}
