package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Informe um nome: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nome := scanner.Text()

	fmt.Println("")
	fmt.Println(ExibirMensagem(nome))
}
