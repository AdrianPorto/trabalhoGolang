package main

import "fmt"

// Crie um programa que apresente o valor da conversão em real (R$) de um valor lido em dólar (US$).
// ➢ O programa deve solicitar o valor da cotação do dólar(REAL).
// ➢ Também a quantidade de dólares disponível com o usuário.
// ➢ Armazenar em memória o valor da conversão antes da apresentação

func main() {
	var cotacaoReal, quantidadeDolar, valorCotado float64

	println("---------------------------------------------------------------")
	println("| Informe a cotação do REAL para dólar:                       |")
	println("---------------------------------------------------------------")
	fmt.Scanln(&cotacaoReal)

	println("---------------------------------------------------------------")
	println("| Informe a quantia em dólar que deseja transformar em reais: |")
	println("---------------------------------------------------------------")
	fmt.Scanln(&quantidadeDolar)

	valorCotado = quantidadeDolar * cotacaoReal

	println("-----------------------------------------")
	fmt.Printf("Valor contato em reais: R$ %.2f\n", valorCotado)
	println("-----------------------------------------")

}
