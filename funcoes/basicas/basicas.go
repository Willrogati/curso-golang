package main

import "fmt"

func f1() { // sem parametro e sem retorno
	fmt.Println("F1")
}

func f2(p1 string, p2 string) { // com parametro sem retorno
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string { //sem paramentro com retorno
	return "F3"
}

func f4(p1, p2 string) string { // com parametro e com retorno
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) { // sem parametro mais com retornos multiplos
	return "Retorno 1", "Retorno2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5", r51, r52)
}
