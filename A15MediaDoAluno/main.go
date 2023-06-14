package main

import "fmt"

func main() {
	fmt.Println("Código para calcular Média do Aluno")

	var nota1, nota2, nota3, nota4 float64

	fmt.Print("Informe o valor da média do primeiro bimestre: ")
	fmt.Scan(&nota1)

	fmt.Print("Informe o valor da média do segundo bimestre: ")
	fmt.Scan(&nota2)

	fmt.Print("Informe o valor da média do terceiro bimestre: ")
	fmt.Scan(&nota3)

	fmt.Print("Informe o valor da média do quarto bimestre: ")
	fmt.Scan(&nota4)

	mediaFinal := (nota1 + nota2 + nota3 + nota4) / 4

	fmt.Print("A média final do aluno foi de: ", mediaFinal)

	if mediaFinal >= 5 {
		fmt.Println("\nAluno Aprovado!!!")
	} else {
		fmt.Println("\nAluno reprovado!!!")
	}
}
