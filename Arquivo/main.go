package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	arquivo, err := os.Open("dados.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro:", err.Error())
		return
	}

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println("O conteúdo da linha é:", linha)
	}

	defer arquivo.Close()
}
