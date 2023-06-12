package main

import (
	"fmt"
	"strings"
)

func main() {
	var texto string

	println("--------------------")
	println("| Digite o texto:  |")
	println("--------------------")
	fmt.Scanln(&texto)

	primeiraLetraNaoRepetida := encontrarPrimeiraLetraNaoRepetida(texto)

	if primeiraLetraNaoRepetida != "" {
		println("-------------------------------------------------------------")
		fmt.Printf("A primeira letra não repetida no texto é: %s\n", primeiraLetraNaoRepetida)
		println("-------------------------------------------------------------")
	} else {
		println("-------------------------------------------------------------")
		println("| Não existem letras que não se repetem no texto informado. |")
		println("-------------------------------------------------------------")
	}
}

func encontrarPrimeiraLetraNaoRepetida(texto string) string {
	letras := make(map[string]int)

	texto = strings.ToLower(texto)

	for _, letra := range texto {
		letras[string(letra)]++
	}

	for _, letra := range texto {
		if letras[string(letra)] == 1 {
			return string(letra)
		}
	}

	return ""
}
