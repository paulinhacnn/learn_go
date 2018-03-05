package main

import (
	"encoding/json"
	"fmt"
	"github.com/paulinhacnn/learn_go/Tratamento_Erros/model"
)

var cache map[string]model.Gene

func main() {

	cache = make(map[string]model.Gene, 0)

	dna := model.Gene{}
	dna.Nome = "Cromossomo 21"
	dna.X = 21
	dna.Y = 22
	dna.SetSequencia("ACTG")
	objJSON, err := json.Marshal(dna)

	// tratamento de erro
	if err != nil {
		fmt.Println("[main]Houve um erro na geração do objeto erro!", err.Error())
		return
	}

	fmt.Printf("Gene é Json: %+v\r\n", string(objJSON))
	fmt.Printf("Gene é: %+v\r\n", dna)
	fmt.Printf("O valor da casa é: %d\r\n", dna.GetSequencia())

	cache["Cromossomo 21"] = dna

	fmt.Println("Encontrou o cromosso 21?")

	dna, achou := cache["Cromossomo 21"]
	if achou {
		fmt.Println("Sim, encontrei: %+v\r\n", dna)
	}

}
