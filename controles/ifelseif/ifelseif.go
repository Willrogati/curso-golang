package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E" // se não atender nenhuma das condicionais acima vai direto
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(5.8))
	fmt.Println(notaParaConceito(4.9))
	fmt.Println(notaParaConceito(2.8))
}
