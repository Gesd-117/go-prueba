package main

import (
	util "arsistema/utilidades"
	"fmt"
)

func main() {
	var (
		altura    float64
		sumatoria float64
		maximo    float64
		minimo    float64
		//promedio float64
	)

	const numeroDePersonas = 6

	avisos := [numeroDePersonas]string{
		"primera",
		"segunda",
		"tercera",
		"cuarta",
		"quinta",
		"sexta",
	}

	util.Encabezado("ingreso de alturas", 60)

	for i := 0; i < numeroDePersonas; i++ {
		altura = util.LeerNumerosDecimales("ingresa " + avisos[i] + " altura (mts)")
		sumatoria += altura

		if i == 0 {
			minimo = altura
			maximo = altura
		}
		if altura > maximo {
			maximo = altura
		}
		if altura < minimo {
			minimo = altura
		}

	}
	util.Linea(60)
	fmt.Printf("La sumatoria es %.2f\n", sumatoria)
	fmt.Printf("La maxima altura es %.2f mts\n", maximo)
	fmt.Printf("La minima altura es %.2f mts\n", minimo)
	util.Linea(60)
}
