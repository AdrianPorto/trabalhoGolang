package main

import "fmt"

func main() {
	// Antes de calcular qualquer informação, precisamos definir os valores das cédulas e moedas
	// Valores das cédulas e moedas
	cedulas := []float64{100, 50, 10, 5, 1}
	moedas := []float64{0.5, 0.1, 0.05, 0.01}

	fmt.Println("Código de Troco")
	fmt.Println("Para calcular o troco a ser recebido, é necessário coletar algumas informações.")

	// Entrada de dados de registro pelo console
	var total, pago float64
	fmt.Print("Digite o valor total a ser pago: ")
	_, err := fmt.Scan(&total)
	if err != nil {
		return
	}
	fmt.Print("Digite o valor efetivamente pago: ")
	_, err = fmt.Scan(&pago)
	if err != nil {
		return
	}

	// Cálculo do troco
	troco := pago - total

	// Impressão das cédulas e moedas do troco
	fmt.Println("Troco: ", troco)

	for _, cedula := range cedulas {
		qtd := int(troco / cedula)
		if qtd > 0 {
			fmt.Printf("Cédulas de R$%.2f: %d\n", cedula, qtd)
			troco -= cedula * float64(qtd)
		}
	}

	for _, moeda := range moedas {
		qtd := int(troco / moeda)
		if qtd > 0 {
			fmt.Printf("Moedas de R$%.2f: %d\n", moeda, qtd)
			troco -= moeda * float64(qtd)
		}
	}
}
