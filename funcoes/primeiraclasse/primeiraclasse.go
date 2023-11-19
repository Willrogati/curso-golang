package main

import "fmt"

var soma = func(a, b int) int { // atribiu uma variavel a uma func anonima
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))
	sub := func(a, b int) int { /* tambem posso criar a func dentro da func main atribuindo
		ela a uma variavel */
		return a - b
	}

	fmt.Println(sub(2, 3))

}
