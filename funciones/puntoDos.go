package funciones

import "fmt"

/**
 * @pre: Ingresar por teclado la respuesta correcta.
 *
 * @post: Si la respuesta es correcta se emite el mensaje "Opcion Correcta" si
 * 		  el mensaje es errado se emite el mensaje "Opcion Incorrecta"
 */
func PuntoDos() {

	for i := 1; i < 5; i++ {
		fmt.Printf("- Opción %d\n", i)
	}

	var opcion int

	fmt.Println("Seleccione el número PAR que sea MAYOR a 2")
	fmt.Print("Opcion: ")
	fmt.Scanln(&opcion)

	if opcion == 4 {
		fmt.Println("Opcion Correcta")
	} else {
		fmt.Println("Opcion Incorrecta")
	}
}
