package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/paulinhacnn/learn_go/Arquivo/model"
	"os"
)

func main() {

	arquivo, err := os.Open("dados.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro:", err.Error())
		return
	}
	//scanner := bufio.NewScanner(arquivo)
	//for scanner.Scan() {
	//		linha := scanner.Text()
	//	fmt.Println("O conteúdo da linha é:", linha)
	//}
	lerCsv := csv.NewReader(arquivo)
	// readall() retorna um slice duplo linha e os elementos
	conteudo, err := lerCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo. Erro:", err.Error())
		return
	}

	arquivoJSON, err := os.Create("medicamento.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro:", err.Error())
		return
	}
	tipos := bufio.NewWriter(arquivoJSON)
	tipos.WriteString("[\r\n")
	for iLinha, linha := range conteudo {
		fmt.Printf("Linha[%d] é %s\r\n", iLinha, linha)
		
	}

	arquivo.Close()
}
