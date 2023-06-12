package main

import "fmt"

func main() {
	var valorHoraAula, horasTrabalhadas, percentualDesconto float64

	println("------------------------------------------------")
	println("| Digite o valor da hora-aula:                 |")
	println("------------------------------------------------")
	fmt.Scanln(&valorHoraAula)

	println("------------------------------------------------")
	println("| Digite o número de horas trabalhadas no mês: |")
	println("------------------------------------------------")
	fmt.Scanln(&horasTrabalhadas)

	println("------------------------------------------------")
	println("| Digite o percentual de desconto do INSS:     |")
	println("------------------------------------------------")
	fmt.Scanln(&percentualDesconto)

	salarioBruto := valorHoraAula * horasTrabalhadas
	desconto := salarioBruto * (percentualDesconto / 100)
	salarioLiquido := salarioBruto - desconto

	println("------------------------------------------------")
	fmt.Printf("Salário Bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Salário Líquido: R$ %.2f\n", salarioLiquido)
	println("------------------------------------------------")
}
