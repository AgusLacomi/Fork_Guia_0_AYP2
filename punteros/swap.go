package punteros

import "fmt"

// permita intercambiar el valor almacenado en dos variables enteras a través de punteros.
func Swap(px, py *int) {
	fmt.Println("punteroX:", *px)
	fmt.Println("punteroY:", *py)
	aux := *px
	*px = *py
	*py = aux
}
