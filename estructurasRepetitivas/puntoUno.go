package estructurasrepetitivas

/**
 * @param numero: numero de tipo entero que se quiere obtener su factorial
 *
 * @pre: El factorial no debe ser negativo
 *
 * @post: retorna el factorial del numero ingresado por parametro
 */
func PuntoUno(numero int) int {

	factorial := 1

	for i := 1; i < int(numero); i++ {
		factorial += factorial * i
	}

	return factorial
}
