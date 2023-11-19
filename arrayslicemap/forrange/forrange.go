package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta quantas posioções precisa no indice

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // usando o _ é ignorado o indice e consigo ler o valor de num
		fmt.Println(num)
	}
}
