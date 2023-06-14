package main

import "fmt"

func main() {
	fmt.Println("Código para Eleição Sindical")
	fmt.Println("Precisamos primeiro analisar a leitura dos votos, para depois efetuar os devidos cálculos")

	var votosValidosA, votosValidosB, votosValidosC, votosNulos, votosBranco int

	fmt.Println("Informe o valor de votos válidos para o candidato A:")
	_, err := fmt.Scan(&votosValidosA)
	if err != nil {
		return
	}

	fmt.Println("Informe o valor de votos válidos para o candidato B:")
	_, err = fmt.Scan(&votosValidosB)
	if err != nil {
		return
	}

	fmt.Println("Informe o valor de votos válidos para o candidato C:")
	_, err = fmt.Scan(&votosValidosC)
	if err != nil {
		return
	}

	fmt.Println("Informe o valor de votos nulos:")
	_, err = fmt.Scan(&votosNulos)
	if err != nil {
		return
	}

	fmt.Println("Informe o valor de votos em branco:")
	_, err = fmt.Scan(&votosBranco)
	if err != nil {
		return
	}

	numTotalEleitores := votosValidosA + votosValidosB + votosValidosC + votosNulos + votosBranco

	// Cálculo do percentual correspondente
	percentualVotosValidos := (100 * (votosValidosA + votosValidosB + votosValidosC)) / numTotalEleitores
	percentualVotosA := (100 * votosValidosA) / numTotalEleitores
	percentualVotosB := (100 * votosValidosB) / numTotalEleitores
	percentualVotosC := (100 * votosValidosC) / numTotalEleitores
	percentualVotosNulos := (100 * votosNulos) / numTotalEleitores
	percentualVotosBranco := (100 * votosBranco) / numTotalEleitores

	fmt.Println("\nRESULTADOS:")
	fmt.Println("Número total de eleitores: $numTotalEleitores")
	fmt.Println("Percentual de votos válidos em relação à quantidade de eleitores: ", percentualVotosValidos)
	fmt.Println("Percentual de votos válidos do candidato A em relação à quantidade de eleitores: ", percentualVotosA)
	fmt.Println("Percentual de votos válidos do candidato B em relação à quantidade de eleitores: ", percentualVotosB)
	fmt.Println("Percentual de votos válidos do candidato C em relação à quantidade de eleitores: ", percentualVotosC)
	fmt.Println("Percentual de votos nulos em relação à quantidade de eleitores: ", percentualVotosNulos)
	fmt.Println("Percentual de votos em branco em relação à quantidade de eleitores: ", percentualVotosBranco)
}
