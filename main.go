package main

import "fmt"

func main() {
	var quantia int
	var nota50 int
	var nota10 int
	var nota5 int
	var nota1 int
	var resto int

	fmt.Println("Bem vindo! Digite a sua quantia: ")
	fmt.Scan(&quantia)

	nota50 = quantia / 50
	resto = quantia % 50
	nota10 = resto / 10
	resto = resto % 10
	nota5 = resto / 5
	resto = resto % 5
	nota1 = resto / 1
	resto = resto % 1

	fmt.Printf("Você precisa de %d nota(s) de 50, %d nota(s) de 10, %d nota(s) de 5, %d nota(s) de 1\n", nota50, nota10, nota5, nota1)

}
