package main

import (
	"guiaCero/arreglos"
	"guiaCero/estructurasrepetitivas"
	"guiaCero/funciones"
)

// By AgusLacomi
func main() {

	/*Funciones*/
	funciones.PuntoUno(3.0, 2.0, 1.0)
	//funciones.PuntoDos()
	funciones.PuntoTres(5, 4, 1, 2)

	/*Estructuras Repetitivas*/
	estructurasrepetitivas.PuntoUno(5)
	estructurasrepetitivas.PuntoDos(3, 20)
	estructurasrepetitivas.PuntoTres(2)

	/*Arreglos*/
	arreglos.PuntoUno([]int{3, 5})
	arreglos.PuntoDos([]int{3, 0}, []int{5, 5})
	arreglos.PuntoTres()
}
