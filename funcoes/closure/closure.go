package main

import "fmt"

func closure() func() {
	x := 10 // a varialvel tem a memoria de onde ele está dentro de que função
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
	/* sendo que mesmo ela (x) sendo utilizado com mesmo nome em outro funcao
	quando a função é utilizada ela carrega o valor que foi utilizado dentro da
	função por isso Closure(fechado) */
}
