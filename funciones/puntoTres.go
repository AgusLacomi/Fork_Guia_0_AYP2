package funciones

import "fmt"

/**
 * @param: Enteros que se quiere saber el maximo y el minimo
 *
 * @pre: Se ingresa una cantidad indefinida de enteros
 *
 * @post: Se devuelve el maximo y el minimo de los valores ingresados por parametro
 */
func PuntoTres(valores ...int) {

	var slice []int = valores[0:]

	maxValue := 0
	minValue := 0

	ordenar(slice)

	maxValue = slice[len(slice)-1]
	minValue = slice[0]

	fmt.Printf("El valor mínimo es %v y el valor máximo es %v", minValue, maxValue)
}

/**
 * @param: Arreglo a ordenar
 *
 * @pre: Se debe ingresar un arreglo
 *
 * @post: Se ordena el arreglo de menos a mayor
 */
func ordenar(slice []int) {

	for i := 1; i < len(slice); i++ {

		posicion := i
		ultimoValor := slice[i]

		for posicion >= 1 && ultimoValor < slice[posicion-1] {

			slice[posicion] = slice[posicion-1]
			posicion--
		}

		slice[posicion] = ultimoValor

	}
}
