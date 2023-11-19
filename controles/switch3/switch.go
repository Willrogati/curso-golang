package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string kk"
	case func():
		return "função"
	default:
		return "não sei"

	}
}

func main() {
	fmt.Println(tipo(3.2))
	fmt.Println(tipo(10))
	fmt.Println(tipo("oi"))
	fmt.Println(tipo(func() {}))
}
