package figuras

import "fmt"

type Rectangulo struct {
	P1, P2 Punto
}

func (r Rectangulo) Mover(incX, incY int) {
	r.P1.Mover(incX, incY)
	r.P2.Mover(incX, incY)
}

func (r Rectangulo) Base() int {
	return r.P2.X - r.P1.X
}

func (r Rectangulo) Altura() int {
	return r.P2.Y - r.P1.Y
}

func (r Rectangulo) Perimetro() int {
	return 2*r.Base() + 2*r.Altura()
}

func (r Rectangulo) Area() int {
	return r.Base() * r.Altura()
}

func (r Rectangulo) ToString() string {
	cadena := fmt.Sprint("Rectangulo: ", r)

	return cadena
}
