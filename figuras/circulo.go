package figuras

//By AgusLacomi
import (
	"fmt"
	"math"
)

type Circulo struct {
	Pto   Punto
	Radio float32
}

func (c Circulo) Area() float32 {
	return c.Radio * c.Radio * math.Pi
}

func (c Circulo) Perimetro() float32 {
	return c.Radio * 2 * math.Pi
}

func (c Circulo) ToString() string {
	cadena := fmt.Sprint("Circulo:", c)

	return cadena
}
