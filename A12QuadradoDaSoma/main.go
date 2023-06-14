package main

import (
	"fmt"
)

func main() {
	fmt.Println("Código para calcular Quadrado das Somas")
	fmt.Println("Informe três valores para que seja feito o quadrado das somas")

	var num1, num2, num3 int

	fmt.Print("Primeiro valor: ")
	// Se ao invés de colocar _ eu colocar uma variável, ele não vai receber o valor do fmt.scan
	// e sim vai receber a quantidade de caracteres ou uma outra possibilidade
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Erro ao ler a variavel num1: ", err)
		return
	}

	fmt.Print("Segundo valor: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Erro ao ler a variavel num2: ", err)
		return
	}

	fmt.Print("Terceiro valor: ")
	_, err = fmt.Scan(&num3)
	if err != nil {
		fmt.Println("Erro ao ler a variavel num3: ", err)
		return
	}

	soma := num1 + num2 + num3
	quadradoSoma := soma * soma

	fmt.Println("A soma é dos três número é: %d\n O quadrado da soma será: %d", soma, quadradoSoma)
}
