package main

import (
	util "arsistema/utilidades"
	"fmt"
	"strings"
)

func main() {

	var (
		cedula    string  = "32.567.086"
		nombre    string  = "Carlos"
		apellido  string  = "Flores"
		edad      int     = 24
		turno     int     = 1     //1 mañana, 2 tarde, 3 noche
		sexo      bool    = false //true mujer, false hombre
		matricula float64 = 1245.66
		altura    float32 = 1.73
	)

	verDatos(cedula, nombre, apellido, edad, turno, sexo, altura, matricula)

}

func verDatos(datos ...any) {

	util.Encabezado("datos del alumno", 60)


	fmt.Println("Cédula:", datos[0])

	nombreCompleto := strings.ToUpper(datos[1].(string)) + " " + strings.ToUpper(datos[2].(string))
	fmt.Printf("Nombre completo: %v \n", nombreCompleto)

	fmt.Printf("Edad: %d años\n", datos[3])

	fmt.Println("Turno:", indentificarTurno(datos[4].(int)))

	fmt.Println("Sexo:", indentificarSexo(datos[5].(bool)))

	fmt.Printf("Altura: %.2f mts\n", datos[6])

	fmt.Printf("Matricula: %.2f\n", datos[7])

}

func indentificarSexo(sexo bool) string {

	if sexo {
		return "Femenino"

	} else {
		return "Masculino"
	}
}

func indentificarTurno(turno int) string {

	switch turno {
	case 1:
		return "Mañana"

	case 2:
		return "Tarde"

	case 3:
		return "Noche"

	default:
		return "Sin turno"
	}
}
