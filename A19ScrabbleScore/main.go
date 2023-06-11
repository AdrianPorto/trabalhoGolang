package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Bem vindo ao jogo ScrabbleScore! Digite uma palavra e eu direi quantos pontos essa palavra possui!")
	reader := bufio.NewReader(os.Stdin)
	palavra, _ := reader.ReadString('\n')
	palavra = strings.ReplaceAll(palavra, " ", "")
	score := 0
	for i := 0; i < len(palavra); i++ {
		switch strings.ToLower(string(palavra[i])) {
		case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			score += 1
		case "d", "g":
			score += 2
		case "b", "c", "m", "p":
			score += 3
		case "f", "h", "v", "w", "y":
			score += 4
		case "k":
			score += 5
		case "j", "x":
			score += 8
		case "q", "z":
			score += 10
		}
	}
	fmt.Printf("A pontuação dessa palavra é: %d\n", score)
}
