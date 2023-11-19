package main

import "fmt"

// não tem operador ternario
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado" ("Esta forma não existe em Go")
	if nota >= 6 {
		return " Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
