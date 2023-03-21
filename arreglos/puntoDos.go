package arreglos

/**
 * @param arregloUno: Arreglo de enteros
 * 		  arregloDos: Arreglo de enteros
 *
 * @pre: Se deben ingresar dos arreglos de enteros necesariamente
 *
 * @post: Se retorna la suma vectorial y el producto escalar entre los 2 arreglos
 */
func PuntoDos(arregloUno, arregloDos []int) ([]int, int) {

	x := arregloUno[0] + arregloDos[0]
	y := arregloUno[1] + arregloDos[1]

	sumaVectorial := []int{x, y}

	x = arregloUno[0] * arregloDos[0]
	y = arregloUno[1] * arregloDos[1]

	productoEscalar := x + y

	return sumaVectorial, productoEscalar
}
