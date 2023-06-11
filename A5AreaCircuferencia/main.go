package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Bem vindo ao programa de calculo de circuferencia , para que eu possa calcular a área de uma" +
		"circuferência , preciso que você insira o valor do raio!")
	scanner := bufio.NewReader(os.Stdin)
	raioStr, _ := scanner.ReadString('\n')
	raio, _ := strconv.ParseFloat(raioStr, 64)
	pi := 3.14159265

	area := pi * (raio * raio)

	fmt.Println("Esse é o resultado do cálculo:", area)
}
