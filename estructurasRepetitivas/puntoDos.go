package estructurasrepetitivas

import "fmt"

func PuntoDos(a, b int) { // a = 5 y b = 2

	var (
		cantidadDeSumas int
		numeroASumar    int
		producto        int
	)

	if a >= b { // Si a es mayor o igual a b se deveria sumar b veces a
		cantidadDeSumas = b
		numeroASumar = a
	} else {
		cantidadDeSumas = a
		numeroASumar = b
	}

	for i := 0; i < cantidadDeSumas; i++ {
		producto += numeroASumar
	}

	fmt.Println("El producto entre", a, "y", b, "es", producto)
}
