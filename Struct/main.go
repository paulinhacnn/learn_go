package main

import "fmt"

type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu apto", 79000}
	fmt.Printf("O Apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		valor: 55,
		X:     22,
	}
	fmt.Printf("A chacara é: %+v\r\n", chacara)

	casa.Nome = "Lar meu pequeno lar"
	casa.valor = 500000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A casa é: %+v\r\n", casa)
}
