package figuras

import "fmt"

type Cuadrado struct {
	Pto  Punto //esquina inferior izquierda
	Lado int
}

func (c *Cuadrado) Mover(incX, incY int) {
	c.Pto.Mover(incX, incY)
}

func (c Cuadrado) Perimetro() int {
	return 4 * c.Lado
}

func (c Cuadrado) Area() int {
	return c.Lado * c.Lado
}

func (c Cuadrado) ToString() string {
	cadena := fmt.Sprint("Cuadrado:", c)

	return cadena
}
