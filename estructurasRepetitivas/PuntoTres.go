package estructurasrepetitivas

import "fmt"

func PuntoTres(numero int) {

	primo := true

	if numero%2 == 0 {
		primo = false
	}

	fmt.Println("Its", primo, "that this number is prime")

}
