package main

import "fmt"
import "github.com/paulinhacnn/learn_go/Funcoes/matematica"

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("O resultado da divisão foi: %d\r\n", resultado)
	resultado, resto := matematica.DivisorResto(14, 3)
	fmt.Printf("O resultado da divisão foi: %d e o resto da divisão foi: %d\r\n", resultado, resto)
}
