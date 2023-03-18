package arreglos

import "fmt"

func PuntoDos(arregloUno, arregloDos []int) {

	x := arregloUno[0] + arregloDos[0]
	y := arregloUno[1] + arregloDos[1]

	sumaVectorial := []int{x, y}

	fmt.Println("La suma vectorial entre", arregloUno, "y", arregloDos, "es:", sumaVectorial)

	x = arregloUno[0] * arregloDos[0]
	y = arregloUno[1] * arregloDos[1]

	productoEscalar := x + y

	fmt.Println("El producto escalar entre", arregloUno, "y", arregloDos, "es:", productoEscalar)

}
