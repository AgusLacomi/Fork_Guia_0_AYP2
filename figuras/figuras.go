package figuras

type Figura interface {
	Area() float32
	Perimetro() float32
	ToString() string
}
