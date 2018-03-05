package model

type Gene struct {
	X         int    `json: "X"`
	Y         int    `json: "Y"`
	Nome      string `json: "nome"`
	sequencia string
}

func (g *Gene) SetSequencia(sequencia string) {
	g.sequencia = sequencia
}
func (g *Gene) GetSequencia() string {
	return g.sequencia
}
