package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {

	comprarTv50 := trab1 && trab2    // E
	comprarTv32 := trab1 != trab2    // ou exclusivo
	comprarSorvete := trab1 || trab2 // ou

	return comprarTv50, comprarTv32, comprarSorvete

}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	// ! a exclamção é operado de negação inverte o valor de saida
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t, inverteTv32: %t", tv50, tv32, sorvete, !sorvete, !tv32)

}
