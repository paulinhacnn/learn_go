package model

type FelisCatus interface {
	Mia() string
}

type CanisLupusFamiliaris interface {
	Late() string
}

type EquusCaballus struct {
	Nome string
}

func (e EquusCaballus) Mia() string {
	return "Miu..."

}
func (e EquusCaballus) Late() string {
	return "Auauau..."

}
