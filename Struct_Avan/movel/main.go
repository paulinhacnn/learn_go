package main

import (
	"github.com/paulinhacnn/learn_go/Struct_Avan/movel"
)

func main() {

	casa := movel.Imovel
	casa.Nome = "Rosa"
	casa.X = 1319
	casa.Y = 84
	casa.SetValor(500000)

	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("O valor da casa é: %+v\r\n", casa.GetValor())
}
