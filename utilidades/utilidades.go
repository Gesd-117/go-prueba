package utilidades

import (
	"fmt"
	"strings"
)

func centrar(texto string, largo int) string {

	relleno := (largo - len(texto)) / 2
	textoCentrado := strings.Repeat(" ", relleno) + texto

	return textoCentrado
}

func Linea(largo int) {
	fmt.Println(strings.Repeat("═", largo))
}

func Encabezado(titulo string, largo int) {
	Linea(largo)
	fmt.Println(centrar(strings.ToUpper(titulo), largo))
	Linea(largo)
}

func LeerNumeroEntero(mensaje string) int{	

var numero int

for  {
	fmt.Print(mensaje + ": ")
	_, err:= fmt.Scanln(&numero)

	if err != nil{
		fmt.Println("Debe escribir un numero entero")

		var limpiar string
		fmt.Scanln(&limpiar)
		continue
	}

	return numero
}
}

func LeerNumerosDecimales(mensaje string) float64{	

var numero float64

for  {
	fmt.Print(mensaje + ": ")
	_, err:= fmt.Scanln(&numero)

	if err != nil{
		fmt.Println("Debe escribir un numero")

		var limpiar string
		fmt.Scanln(&limpiar)
		continue
	}

	return numero
}

	
}
