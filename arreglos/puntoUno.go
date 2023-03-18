package arreglos

import "fmt"

func PuntoUno(arreglo []int) {

	var sumaElementos int

	for i := 0; i < len(arreglo); i++ {
		sumaElementos += arreglo[i]
	}

	fmt.Println("La suma de todos los elementos del arreglo ingresado es", sumaElementos)

}
