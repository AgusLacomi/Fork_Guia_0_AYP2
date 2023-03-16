package funciones

import "fmt"

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
