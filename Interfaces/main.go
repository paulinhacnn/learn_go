package main

import (
	"fmt"
	"github.com/paulinhacnn/learn_go/Interfaces/model"
)

func main() {

	lilo := model.EquusCaballus{}
	lilo.Nome = "Lilo Cavalo"
	SomGato(lilo)

}
func SomGato(g model.FelisCatus) {
	fmt.Println(g.Mia())
}
func SomCao(p model.CanisLupusFamiliaris) {
	fmt.Println(p.Late())
}
