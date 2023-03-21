package figuras

import "fmt"

type Circulo struct {
	Pto   Punto
	Radio int
}

func (c Circulo) Area() int {
	return c.Radio * c.Radio
}

func (c Circulo) Perimetro() int {
	return c.Radio * 2
}

func (c Circulo) ToString() string {
	cadena := fmt.Sprint("Circulo:", c)

	return cadena
}
