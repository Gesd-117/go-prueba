package main

import (
	util "arsistema/utilidades"
	"fmt"
)

func main() {

	util.Encabezado("hola mundo", 36)

	sumar()
}

func sumar() {
	var (
		numero1 int
		numero2 int
		numero3 int
		numero4 int
		resultado int
	)

	util.Encabezado("suma de numeros", 60)

	numero1 = util.LeerNumeroEntero("Ingresa el primer numero")
	numero2 = util.LeerNumeroEntero("Ingresa el segundo numero")
	numero3 = util.LeerNumeroEntero("Ingresa el tercer numero")
	numero4 = util.LeerNumeroEntero("Ingresa el cuarto numero")
	util.Linea(60)

	resultado = sumaDeNumeros(numero1,numero2,numero3,numero4)

	fmt.Printf("El resultado de la suma es %d\n", resultado)
	util.Linea(60)
}

func sumaDeNumeros(numeros ...int) int{

	var sumatoria int

	for _, numero := range numeros {
		sumatoria += numero
	}

	return sumatoria
	
}
