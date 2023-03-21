package arreglos

import "fmt"

func PuntoTres(arregloA, arregloB []int) []int {

	union := arregloA

	for _, numero := range arregloA {

		for i := 0; i < len(arregloB); i++ {

			if numero != arregloB[i] {
				union = append(union, arregloB[i])
			}
		}
	}

	union = ordenar(union)

	fmt.Println("despues de ordenar:", union)

	return union
}

/**
 * @param arreglo: Arreglo a ordenar
 *
 * @pre: Se debe ingresar un arreglo
 *
 * @post: Se ordena el arreglo de menor a mayor y se eliminan numeros del arreglo repetidos
 */
func ordenar(arreglo []int) []int {

	for i := 1; i < len(arreglo); i++ {

		posicion := i
		ultimoValor := arreglo[i]

		for posicion >= 1 && ultimoValor < arreglo[posicion-1] {
			arreglo[posicion] = arreglo[posicion-1]
			posicion--
		}

		arreglo[posicion] = ultimoValor
	}

	arreglo = eliminarDuplicados(arreglo)

	return arreglo
}

func eliminarDuplicados(arreglo []int) []int {

	for i := 0; i < len(arreglo)-1; i++ {

		if arreglo[i] == arreglo[i+1] {
			arreglo[i+1] = 0
		}
	}

}
