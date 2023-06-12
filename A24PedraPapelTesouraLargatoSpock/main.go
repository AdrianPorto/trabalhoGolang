package main

import (
	"fmt"
	"strings"
)

//➢ Pedra-papel-tesoura-lagarto-Spock é uma adaptação do JoKenPo na série The Big Bang Theory.
//➢ Crie um programa que seja um juiz de Jokenpo (Sheldon Cooper) que dada a jogada dos dois jogadores informa o
//resultado da partida.
//➢ As regras são as seguintes:
//○ Tesoura corta papel
//○ Papel cobre pedra
//○ Pedra esmaga lagarto
//○ Lagarto envenena Spock
//○ Spock esmaga (ou derrete) tesoura
//○ Tesoura decapita lagarto
//○ Lagarto come papel
//○ Papel refuta Spock
//○ Spock vaporiza pedra
//○ Pedra amassa tesoura

func main() {
	var jogador1, jogador2 string

	println("---------------------------------------------------------------------")
	println("| Jogador 1, escolha entre pedra, papel, tesoura, lagarto ou Spock: |")
	println("---------------------------------------------------------------------")
	fmt.Scanln(&jogador1)

	println("---------------------------------------------------------------------")
	println("| Jogador 2, escolha entre pedra, papel, tesoura, lagarto ou Spock: |")
	println("---------------------------------------------------------------------")
	fmt.Scanln(&jogador2)

	resultado := verificarResultado(jogador1, jogador2)

	println("---------------------------------------------------------------------")
	fmt.Println(resultado)
	println("---------------------------------------------------------------------")
}

func verificarResultado(jogador1, jogador2 string) string {
	regras := map[string]string{
		"tesoura-papel":   "Tesoura corta papel",
		"papel-pedra":     "Papel cobre pedra",
		"pedra-lagarto":   "Pedra esmaga lagarto",
		"lagarto-spock":   "Lagarto envenena Spock",
		"spock-tesoura":   "Spock esmaga tesoura",
		"tesoura-lagarto": "Tesoura decapita lagarto",
		"lagarto-papel":   "Lagarto come papel",
		"papel-spock":     "Papel refuta Spock",
		"spock-pedra":     "Spock vaporiza pedra",
		"pedra-tesoura":   "Pedra amassa tesoura",
	}

	jogada1 := strings.ToLower(jogador1)
	jogada2 := strings.ToLower(jogador2)
	jogada := jogada1 + "-" + jogada2

	if regra, ok := regras[jogada]; ok {
		return fmt.Sprintf("Jogador 1 escolheu %s\nJogador 2 escolheu %s\n%s\nJogador 1 venceu!", jogador1, jogador2, regra)
	} else if regra, ok := regras[jogada2+"-"+jogada1]; ok {
		return fmt.Sprintf("Jogador 1 escolheu %s\nJogador 2 escolheu %s\n%s\nJogador 2 venceu!", jogador1, jogador2, regra)
	} else {
		return fmt.Sprintf("Jogador 1 escolheu %s\nJogador 2 escolheu %s\nNenhum jogador venceu, é um empate!", jogador1, jogador2)
	}
}
