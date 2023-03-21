package arreglos

func PuntoTres(arregloA, arregloB []int) ([]int, []int) {

	unionDeArreglos := unionDeConjuntos(arregloA, arregloB)

	interseccionDeArreglos := interseccionDeConjuntos(arregloA, arregloB)

	retornoUnion := eliminarRepetidos(unionDeArreglos)
	retornoInterseccion := eliminarRepetidos(interseccionDeArreglos)

	return retornoUnion, retornoInterseccion
}

func unionDeConjuntos(arregloA []int, arregloB []int) []int {

	unionDeArreglos := arregloA[0:]

	unionDeArreglos = append(unionDeArreglos, arregloB...)

	ordenar(unionDeArreglos)

	return unionDeArreglos
}

func interseccionDeConjuntos(arregloA, arregloB []int) []int {

	var interseccionDeArreglos []int

	for _, numeroA := range arregloA {
		for _, numeroB := range arregloB {

			if numeroA == numeroB {
				interseccionDeArreglos = append(interseccionDeArreglos, numeroA)
			}
		}
	}

	ordenar(interseccionDeArreglos)

	return interseccionDeArreglos
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

	return arreglo
}

/**
 * @param arr: el arreglo original con elementos repetidos
 *
 * @pre: arr debe ser un arreglo de enteros
 *
 * @post: devuelve un nuevo arreglo sin elementos repetidos arreglos
 */
func eliminarRepetidos(arreglo []int) []int {

	Clave := make(map[int]bool)

	lista := []int{}

	for _, entrada := range arreglo {
		if _, value := Clave[entrada]; !value {
			Clave[entrada] = true
			lista = append(lista, entrada)
		}
	}

	return lista
}
