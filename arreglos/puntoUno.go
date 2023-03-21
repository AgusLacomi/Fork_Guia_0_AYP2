package arreglos

/**
 *@param arreglo: Arreglo de tipo entero que contiene los elementos a sumar.
 *
 *@pre: Se ingresa un arreglo de tipo entero con los elementos a sumar.
 *
 * @post: Retorna la suma de todos los elementos del arreglo.
 */
func PuntoUno(arreglo []int) int {

	var sumaElementos int

	for i := 0; i < len(arreglo); i++ {
		sumaElementos += arreglo[i]
	}

	return sumaElementos
}
