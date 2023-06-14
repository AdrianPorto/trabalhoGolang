package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declaração do array e leitura dos elementos
	var numeros [12]int

	fmt.Println("Digite 12 números inteiros:")
	for i := 0; i < len(numeros); i++ {
		fmt.Printf("Número %d: ", i+1)
		_, err := fmt.Scan(&numeros[i])
		if err != nil {
			return
		}
	}

	// Método sort importado, para poder ordenar de forma decrescente
	sort.Sort(sort.Reverse(sort.IntSlice(numeros[:])))

	// Método for para imprimir os elementos na ordem decrescente
	fmt.Println("Elementos em ordem decrescente:")
	for _, numero := range numeros {
		fmt.Println(numero)
	}
}
