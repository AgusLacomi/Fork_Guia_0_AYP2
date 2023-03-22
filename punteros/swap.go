package punteros

// permita intercambiar el valor almacenado en dos variables enteras a través de punteros.
func Swap(px, py *int) {
	aux := *px
	*px = *py
	*py = aux
}
