package main

import (
	"fmt"
	"guiaCero/arreglos"
	"guiaCero/estructurasrepetitivas"
	"guiaCero/funciones"
)

// By AgusLacomi
func main() {

	/*Funciones*/
	funciones.PuntoUno(3.0, 2.0, 1.0)
	funciones.PuntoDos()
	minValue, maxValue := funciones.PuntoTres(5, 4, 1, 2)
	fmt.Printf("El valor mínimo es %v y el valor máximo es %v\n", minValue, maxValue)

	/*Estructuras Repetitivas*/
	numero := 5
	factorial := estructurasrepetitivas.PuntoUno(numero)
	fmt.Printf("El factorial de %v es %v\n", numero, factorial)

	a := 5
	b := 2
	producto := estructurasrepetitivas.PuntoDos(a, b)
	fmt.Println("El producto entre", a, "y", b, "es", producto)

	primo := estructurasrepetitivas.PuntoTres(3)
	fmt.Println("Its", primo, "that this number is prime")

	/*Arreglos */
	sumaElementos := arreglos.PuntoUno([]int{-3, 5})
	fmt.Println("La suma de todos los elementos del arreglo ingresado es", sumaElementos)

	arregloUno := []int{3, 0}
	arregloDos := []int{5, 5}
	sumaVectorial, productoEscalar := arreglos.PuntoDos(arregloUno, arregloDos)
	fmt.Println("La suma vectorial entre", arregloUno, "y", arregloDos, "es:", sumaVectorial)
	fmt.Println("El producto escalar entre", arregloUno, "y", arregloDos, "es:", productoEscalar)

	arregloA := []int{3, 2, 1}
	arregloB := []int{3, 5, 4, 6}
	union := arreglos.PuntoTres(arregloA, arregloB)
	fmt.Println("La union entre los dos arreglos es:", union)
}
