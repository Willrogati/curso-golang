package main

import "fmt"

// os ... significa função variatica
func media(numeros ...float64) float64 { // a qantidade de numeros inseridos no caso notas é variavel
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros)) // preciso converter a quantidade de numeros inseridos de Int para Float para realizar o caulculo
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
	fmt.Println(numeros)
}
