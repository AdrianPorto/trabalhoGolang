package main

import "fmt"

func main() {
	var cotacaoDolar, quantidadeReais, valorCotado float64

	println("---------------------------------------------------------------")
	println("| Informe a cotação do DÓLAR para real:                       |")
	println("---------------------------------------------------------------")
	fmt.Scanln(&cotacaoDolar)

	println("---------------------------------------------------------------")
	println("| Informe a quantia em reais que deseja transformar em dólar: |")
	println("---------------------------------------------------------------")
	fmt.Scanln(&quantidadeReais)

	valorCotado = quantidadeReais / cotacaoDolar

	println("-----------------------------------------")
	fmt.Printf("Valor contato em dólar: US$ %.2f\n", valorCotado)
	println("-----------------------------------------")
}
