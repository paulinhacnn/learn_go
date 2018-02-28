package movel

type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}
func (i *Imovel) GetValor() int {
	return i.valor
}
