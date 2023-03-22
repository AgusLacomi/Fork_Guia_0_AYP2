package figuras

import "fmt"

type Punto struct {
	X float32
	Y float32
}

func (p *Punto) ToString() string {
	cadena := fmt.Sprintf("(%v, %v)", p.X, p.Y)
	return cadena
}

func (p *Punto) Mover(x, y float32) {
	p.X += x
	p.Y += y
}
