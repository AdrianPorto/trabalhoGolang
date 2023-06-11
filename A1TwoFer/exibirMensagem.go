package main

import "fmt"

func ExibirMensagem(nome string) string {
	if nome == "" {
		return "Um para você, um para mim."
	} else {
		return fmt.Sprintf("Um para %s, um para mim.", nome)
	}
}
