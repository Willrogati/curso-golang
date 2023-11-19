package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func soma(a, b int) int {
	return a + b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(soma, 3, 4)
	fmt.Println(resultado)
	resultado2 := exec(multiplicacao, 3, 5)
	fmt.Println(resultado2)
}
