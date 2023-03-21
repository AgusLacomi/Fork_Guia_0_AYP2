package figuras

import "fmt"

type Punto struct {
	X int
	Y int
}

func (p *Punto) ToString() string {
	cadena := fmt.Sprintf("(%v, %v)", p.X, p.Y)
	return cadena
}

func (p *Punto) Mover(x, y int) {
	p.X += x
	p.Y += y
}
