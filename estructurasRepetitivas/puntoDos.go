package estructurasrepetitivas

/**
 * @param a: numero a de tipo entero
 * 		  b: numero b de tipo entero
 *
 * @pre: Se ingresan dos numeros enteros de los que se quiere obtener el valor del producto.
 *
 * @post: retorna el producto entre los dos numeros mediante una suma sucesiva
 *
 * [20/3/23] Ahora el producto se puede hacer si hay alguno o ambos de los numero ingresados negativos
 */
func PuntoDos(a, b int) int { // a = -5 y b = -2

	var (
		cantidadDeSumas  int
		numeroASumar     int
		producto         int
		productoNegativo bool
	)

	a, b, productoNegativo = parametroNegativo(a, b, productoNegativo)

	if a >= b { // Si a es mayor o igual a b se deveria sumar b veces a
		cantidadDeSumas = b
		numeroASumar = a
	} else {
		cantidadDeSumas = a
		numeroASumar = b
	}

	for i := 0; i < cantidadDeSumas; i++ {
		producto += numeroASumar
	}

	if productoNegativo {
		producto = producto * -1
	}

	return producto
}

/**
 * @param a: numero a de tipo entero
 * 		  b: numero b de tipo entero
 *        productoNegativo: variable de tipo booleana
 *
 * @pre: Se ingresa por parametro los 2 numero a multiplicar y una variable booleana
 *       que indicara mediante un XOR si hay un valor negativo.
 *
 * @post: En caso de haber un solo negativo se pasa a positivo el valor negativo y se retorna ambos enteros
 * 		  y la variable @param productoNegativo en true.
 *		  En caso de haber dos negativos se pasan ambos valores a positivo y se retorna ambos enteros y la variable
 *		  @param productoNegativo en false
 */
func parametroNegativo(a int, b int, productoNegativo bool) (int, int, bool) {

	if a < 0 || b < 0 {

		if !(a < 0 && b < 0) {
			productoNegativo = true
		}

		if a < 0 {
			a *= -1
		}
		if b < 0 {
			b *= -1
		}
	}
	return a, b, productoNegativo
}
