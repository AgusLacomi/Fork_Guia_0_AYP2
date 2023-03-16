package main

import (
	"fmt"
)

// By AgusLacomi
func main() {

	resultado := CalculoPolinomio(-6.0, 5.0, 4.0, 3.0, 2.0, 1.0)

	for _, coeficiente := range resultado {
		fmt.Print(coeficiente)
	}

	fmt.Print("")

}

//By AgusLacomi
/*
	https://www.digitalocean.com/community/tutorials/an-introduction-to-the-strings-package-in-go-es
*/
func CalculoPolinomio(coeficientes ...float32) []string {

	arreglo := make([]string, len(coeficientes)) // make() crea un nuevo slice de strings
	// con la cantidad de elementos igual a la cantidad de variables ingresadas

	for i := 0; i < len(coeficientes); i++ {

		if i > 1 {
			arreglo[i] = fmt.Sprintf(" %+.1f X**%v ", coeficientes[i], i) // (%+ .1f) Con este ejemplo el input sería "-3.1416" y el output sería "- 3.1 ""
		} else if i == 1 {
			arreglo[i] = fmt.Sprintf(" %+.1f X ", coeficientes[i])
		} else {
			arreglo[i] = fmt.Sprintf(" %.1f ", coeficientes[i])
		}
	}

	return arreglo
}
