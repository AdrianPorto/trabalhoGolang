package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string

	println("------------------------------")
	println("| Digite uma frase:          |")
	println("------------------------------")
	fmt.Scanln(&frase)

	if ehPangrama(frase) {
		println("------------------------------")
		println("| A frase é um pangrama!     |")
		println("------------------------------")
	} else {
		println("------------------------------")
		println("| A frase não é um pangrama. |")
		println("------------------------------")
	}
}

func ehPangrama(frase string) bool {
	frase = strings.ToLower(frase)

	letras := make(map[rune]bool)

	for _, char := range frase {
		if char >= 'a' && char <= 'z' {
			letras[char] = true
		}
	}

	return len(letras) == 26
}
