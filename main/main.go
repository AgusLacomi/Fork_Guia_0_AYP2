package main

import (
	"fmt"
)

// By AgusLacomi
func main() {

	//PuntoUno(-6.0, 5.0, 4.0, 3.0, 2.0, 1.0)

	PuntoDos()

}

/**
 * @param coeficientes: Coeficientes de tipo float
 * @pre: Se ingresa una cantidad indefinida de coeficientes necesariamente de tipo float
 *		para armar el polinomio.
 *
 * @post: imprime por pantalla el polinomio formado con los parametros ingresados.
 */
func PuntoUno(coeficientes ...float32) {

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

	for _, coeficiente := range arreglo {
		fmt.Print(coeficiente)
	}
}

/**
 *
 */
func PuntoDos() {

	for i := 1; i < 5; i++ {
		fmt.Printf("- Opción %d\n", i)
	}

	var opcion int

	fmt.Println("Seleccione el número PAR que sea MAYOR a 2")
	fmt.Print("Opcion: ")
	fmt.Scanln(&opcion)

	if opcion == 4 {
		fmt.Println("Opcion Correcta")
	} else {
		fmt.Println("Opcion Incorrecta")
	}
}
