package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[963852741] = "Pedro"
	aprovados[543215432] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[543215432])
	delete(aprovados, 543215432)
	fmt.Println(aprovados[543215432])
}
