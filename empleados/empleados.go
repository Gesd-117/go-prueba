package main

import (
	util "arsistema/utilidades"
	"fmt"
	"strings"
)

func main() {
	empleado := map[string]any{
	"cedula": "32.567.086",
	"nombre": "carlos",
	"apellido": "rodriguez",
	"edad": 24,
	"sueldo": 1245.60,
	}

	util.Linea(50)
	fmt.Printf("Cédula: %v\n", empleado["cedula"])

	mayusNomb := strings.ToUpper(empleado["nombre"].(string))
	mayusApell := strings.ToUpper(empleado["apellido"].(string)) 
	nombreCompleto := mayusNomb + " " + mayusApell 
	fmt.Printf("Nombre completo: %v\n", nombreCompleto)

	fmt.Printf("Edad: %v años\n", empleado["edad"])
	fmt.Printf("Sueldo: %v$\n", empleado["sueldo"])


	util.Encabezado("datos del empleado", 50)
	for clave, valor := range empleado {
		fmt.Printf("%v: %v \n", strings.ToUpper(clave), verDatos(clave,valor))
	}
	util.Linea(50)
}

func verDatos(clave string, valor any) any {

	switch clave {
	case "nombre": return  strings.ToUpper(valor.(string))
	case "apellido": return  strings.ToUpper(valor.(string))
	default: return valor 
	}
}