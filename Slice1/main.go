package main

import (
	"fmt"
)

func main() {
	// slice suporta tipo primitivo e string
	// no slice não tem tamanho fixo como no array

	var numbers []int
	fmt.Println(numbers, len(numbers), cap(numbers))
	//inicializando um slice
	numbers = make([]int, 10)
	fmt.Println(numbers, len(numbers), cap(numbers))
	nomes := []string{"Mew"}
	fmt.Println(nomes, len(nomes), cap(nomes))
	//adicionar itens no array
	nomes = append(nomes, "Fourier", "Issac", "Spyke")
	fmt.Println(nomes, len(nomes), cap(nomes))
	// alterar um item do slice
	nomes[2] = "Isac"
	fmt.Println(nomes, len(nomes), cap(nomes))
	numeros := make([]string, 4)
	numeros[0] = "um"
	numeros[1] = "dois"
	numeros[2] = "três"
	numeros[3] = "quatro"
	fmt.Println(numeros, len(numeros), cap(numeros))
	for indice, num := range numeros {
		fmt.Println("Numeros[%d] = %s\r\n", indice, num)
	}

	numEle := numeros[1:3]
	fmt.Println(numEle)
	numEle1 := numeros[:2]
	fmt.Println(numEle1)
	numEle2 := numeros[2:]
	fmt.Println(numEle2)
	numEle3 := numeros[len(numeros)-2:]
	fmt.Println(numEle3)

	indiceRetirar := 2
	//slice temporário
	t := numeros[:indiceRetirar]
	t = append(t, numeros[indiceRetirar+1:]...)
	copy(numeros, t)
	fmt.Println(numeros)
}
