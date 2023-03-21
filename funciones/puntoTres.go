package funciones

/**
 * @param: Enteros que se quiere saber el maximo y el minimo
 *
 * @pre: Se ingresa una cantidad indefinida de enteros
 *
 * @post: Se devuelve el maximo y el minimo de los valores ingresados por parametro
 *
 * [20/3/21] No hace falta ordenar un arreglo para devolver el maximo y el minimo
 * asi se ahorra mas memoria
 */
func PuntoTres(valores ...int) (int, int) {

	maxValue := valores[0]
	minValue := valores[0]

	for _, posicion := range valores[1:] {
		if posicion > maxValue {
			maxValue = posicion
		}
		if posicion < minValue {
			minValue = posicion
		}
	}

	return minValue, maxValue

}
