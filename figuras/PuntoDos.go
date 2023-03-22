package figuras

import (
	"fmt"
	"strings"
)

func PuntoDos() {

	var (
		opcionSeleccionada string
		arrOpciones        []string
	)

	fmt.Print("\nSeleccione 5 figuras geométricas\n- a. Rectangulo\n- b. Cuadrado\n- c. Circulo\n")

	for i := 0; i < 5; i++ {

		fmt.Scanln(&opcionSeleccionada)
		opcionSeleccionada = strings.ToUpper(opcionSeleccionada)
		fmt.Println("A seleccionado la opcion:", opcionSeleccionada)

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
}
