package main

import "fmt"

func main() {
	var inicioIntervalo, fimIntervalo int

	println("----------------------------------")
	println("| Informe o início do intervalo: |")
	println("----------------------------------")
	fmt.Scanln(&inicioIntervalo)

	println("----------------------------------")
	println("| Informe o fim do intervalo:    |")
	println("----------------------------------")
	fmt.Scanln(&fimIntervalo)

	somaPares := 0

	for i := inicioIntervalo; i <= fimIntervalo; i++ {
		if i%2 == 0 {
			somaPares += i
		}
	}

	println("------------------------------------------------------------------")
	fmt.Printf(" A soma dos números pares no intervalo [%d a %d] é: %d\n", inicioIntervalo, fimIntervalo, somaPares)
	println("------------------------------------------------------------------")
}
