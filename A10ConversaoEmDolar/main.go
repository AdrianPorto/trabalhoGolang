package main

import "fmt"

// Crie um programa que apresente o valor da conversão em dólar (US$) de um valor lido em real (R$)
// O programa deve solicitar o valor da cotação do dólar.
// ➢ Também a quantidade de reais disponível com o usuário.
// ➢ Armazenar em memória o valor da conversão antes da apresentação.

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
