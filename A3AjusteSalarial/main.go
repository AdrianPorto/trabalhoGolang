package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	var salario float64
	var percentual float64

	fmt.Println("Bem vindo ao programa de ajuste salarial , por favor informe o seu salário atual: ")
	fmt.Scan(&salario)

	fmt.Println("Agora informe o valor percentual para reajuste do salário , (sem % , ex: 20): ")
	fmt.Scan(&percentual)

	reajuste := salario * percentual / 100
	novoSalario := salario + reajuste

	fmt.Printf("Novo salário: %.2f\n", novoSalario)
}
