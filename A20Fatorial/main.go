package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Olá bem vindo ao programa para calcular o fatorial , entre com um número para calcular seu fatorial:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	numero, _ := strconv.Atoi(input)

	fatorial := calcularFatorial(numero)

	fmt.Printf("O fatorial de %d é %d\n", numero, fatorial)
}
func calcularFatorial(n int) int {
	if n <= 1 {
		return 1
	}

	fatorial := 1
	for i := 2; i <= n; i++ {
		fatorial *= i
	}

	return fatorial
}
