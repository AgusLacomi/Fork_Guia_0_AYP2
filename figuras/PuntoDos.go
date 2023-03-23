package figuras

import (
	"fmt"
	"strings"
)

func PuntoDos() {

	var (
		opcionSeleccionada string
		arrOpciones        []string
		arrFiguras         = [5]Figura{}
	)

	fmt.Println("\nSeleccione 5 figuras geométricas entre las siguientes opciones:\n	- a. Rectangulo\n	- b. Cuadrado\n	- c. Circulo")

	for i := 0; i < 5; i++ {

		fmt.Scanln(&opcionSeleccionada)
		opcionSeleccionada = strings.ToUpper(opcionSeleccionada)
		fmt.Println("Usted a seleccionado la opcion:", opcionSeleccionada)

		switch {
		case opcionSeleccionada == "A":
			arrOpciones = append(arrOpciones, opcionSeleccionada)
		case opcionSeleccionada == "B":
			arrOpciones = append(arrOpciones, opcionSeleccionada)
		case opcionSeleccionada == "C":
			arrOpciones = append(arrOpciones, opcionSeleccionada)
		default:
			i--
			fmt.Println("Usted no a seleccionado una opción válida")

		}
	}

	for i := 0; i < 5; i++ {
		switch {
		case arrOpciones[i] == "A":
			arrFiguras[i] = crearRectangulo(i)
		case arrOpciones[i] == "B":
			arrFiguras[i] = crearCuadrado(i)
		case arrOpciones[i] == "C":
			arrFiguras[i] = crearCirculo(i)
		}
	}

	for _, figura := range arrFiguras {
		mostrar(figura)
	}
	fmt.Println("")
}

func crearCirculo(i int) Figura {
	fmt.Println("Ingrese los parametros de la figura", (i + 1), "[CIRCULO]\nPunto: ")
	var puntoScan int
	var puntoUno Punto
	fmt.Scan(&puntoScan)
	puntoUno.X = float32(puntoScan)
	fmt.Scan(&puntoScan)
	puntoUno.Y = float32(puntoScan)
	fmt.Println("Ingrese los parametros de la figura", (i + 1), "[CIRCULO]\nRadio: ")
	fmt.Scan(&puntoScan)
	return Circulo{Pto: puntoUno, Radio: float32(puntoScan)}
}

func crearCuadrado(i int) Figura {
	fmt.Println("Ingrese los parametros de la figura ", (i + 1), "[CUADRADO]\nPunto: ")
	var puntoScan int
	var puntoUno Punto
	fmt.Scan(&puntoScan)
	puntoUno.X = float32(puntoScan)
	fmt.Scan(&puntoScan)
	puntoUno.Y = float32(puntoScan)
	fmt.Println("Ingrese los parametros de la figura ", (i + 1), "[CUADRADO]\nTamaño de lado: ")
	fmt.Scan(&puntoScan)
	return Cuadrado{Pto: puntoUno, Lado: float32(puntoScan)}
}

func crearRectangulo(i int) Figura {
	fmt.Println("Ingrese los parametros de la figura", (i + 1), "[RECTANGULO]\nPunto 1: ")
	var puntoScan int
	var puntoUno Punto
	fmt.Scan(&puntoScan)
	puntoUno.X = float32(puntoScan)
	fmt.Scan(&puntoScan)
	puntoUno.Y = float32(puntoScan)
	fmt.Println("Ingrese los parametros de la figura", (i + 1), "[RECTANGULO]\nPunto 2: ")
	var puntoDos Punto
	fmt.Scan(&puntoScan)
	puntoDos.X = float32(puntoScan)
	fmt.Scan(&puntoScan)
	puntoDos.Y = float32(puntoScan)
	return Rectangulo{P1: puntoUno, P2: puntoDos}
}

func mostrar(f Figura) {
	fmt.Println(f.ToString())
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
	fmt.Println()
}
