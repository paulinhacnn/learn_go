package main

import (
	"fmt"
	"github.com/paulinhacnn/learn_go/pacotes/prefixo"
	"github.com/paulinhacnn/learn_go/pacotes/operadora"
)
// NomeDoUsuario é o nome do usuário no sistema
var NomeDoUsuario = "Paula"


func main() {
	fmt.Printf("Nome do Usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TestePrefixo)		
}