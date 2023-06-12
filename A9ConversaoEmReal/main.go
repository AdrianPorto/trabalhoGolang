package main

import "fmt"

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
