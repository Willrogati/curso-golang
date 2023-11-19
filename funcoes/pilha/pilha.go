package main

import "runtime/debug"

func f3() {
	debug.PrintStack() // imprimi a pilha do programa
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
