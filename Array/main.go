package main

import "fmt"

func main() {

	var dados [3]int
	dados[0] = 3
	dados[1] = 2
	dados[2] = 1
	fmt.Println("Qual a capacidade desse array", len(dados))

	pessoas := [4]string{"Paula", "Mew", "Fourier", "Issac"}
	fmt.Println("Nomes: \n\r%v\r\n", pessoas)

	lgg := [...]string{"Go", "Python", "C", "C++", "C#", "JavaScript"}
	for indice, nomes := range lgg {
		fmt.Println("Linguagens de Programação[%d] é %s\n\r", indice, nomes)
	}

}
