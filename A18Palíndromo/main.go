package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Bem vindo ao verificador de Palíndromos. Escreva uma palavra, frase ou número e eu irei informar se é um Palíndromo.")
	reader := bufio.NewReader(os.Stdin)
	palavra, _ := reader.ReadString('\n')
	palavra = strings.TrimSpace(palavra)

	if strings.ToLower(palavra) == inverterPalavra(strings.ToLower(palavra)) {
		fmt.Println("É um Palíndromo!")
	} else {
		fmt.Println("Não é um Palíndromo!")
	}
}
func inverterPalavra(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
