package main

import "fmt"

func main() {

	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println(" O valo de x é " + x) deste jeito não funciona pois x é float e não uma string
	xs := fmt.Sprint(x) // retorna o valor de x em string na variavel xs
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	// %f onde f significa o tipo da varial que vou exibir poderia ser s de string
	fmt.Printf("O valor de x é %.2f.", x) //exibe o valor cortando a casa decimal

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
