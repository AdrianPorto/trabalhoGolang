package main

import "fmt"

func main() {
	var numero int

	println("---------------------------")
	println("| Digite um número:       |")
	println("---------------------------")
	fmt.Scanln(&numero)

	println("---------------------------")
	println("| Sequência de Fibonacci: |")
	for i := 0; i <= numero; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
