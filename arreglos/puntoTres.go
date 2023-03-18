package arreglos

func PuntoTres(arregloA, arregloB []int) {

	ordenar(arregloA)
	ordenar(arregloB)

	Union := arregloA

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

func buscar(vector []int, buscado int) bool {

	var (
		inicio     int
		ultimo     = len(vector) - 1
		encontrado bool
	)

	for inicio <= ultimo {

		posAnalisis := (ultimo + inicio) / 2

		if vector[posAnalisis] == buscado {

			encontrado = true

		} else if vector[posAnalisis] < buscado {

			inicio = posAnalisis

		} else if vector[posAnalisis] > buscado {

			ultimo = posAnalisis
		}

	}
	return encontrado
}
