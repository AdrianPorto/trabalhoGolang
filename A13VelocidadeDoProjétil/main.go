package main

import "fmt"

func main() {
	fmt.Println("Código para calcular velocidade de um projétil")
	var dist float64
	var time int

	fmt.Println("Informe distância que será percorrida pelo projétil: ")
	_, err := fmt.Scan(&dist)
	if err != nil {
		return
	}

	fmt.Println("Informe o tempo que demorou para ser percorrido: ")
	_, err = fmt.Scan(&time)
	if err != nil {
		return
	}

	vel := (dist * 100) / (float64(time) * 60)

	fmt.Println("A velocidade percorrida do projeto foi de: ", vel)
}
