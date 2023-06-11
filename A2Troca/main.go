package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Informe um valor para a variável A: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	A := scanner.Text()
	fmt.Println("Informe um valor para a variável B: ")
	scanner.Scan()
	B := scanner.Text()
	copyA := A

	A = B
	B = copyA

	fmt.Printf("O valor da variável A ficou %v e da variável B ficou %v", A, B)
}
