package main

import (
	util "arsistema/utilidades"
	"fmt"
	"strconv"
)

func main() {

	var (
		numero int
		inicio int
		fin    int
		titulo string
	)

	util.Encabezado("generador de tablas de multiplicar", 60)
	numero = util.LeerNumeroEntero("ingresa tu numero")
	titulo = "tabla de multiplicar del " + strconv.Itoa(numero)
	util.Encabezado(titulo, 60)
	generarTabla(numero)

	util.Encabezado("generar rango de tablas", 60)
	inicio = util.LeerNumeroEntero("ingresa numero de inicio")
	fin = util.LeerNumeroEntero("ingresa numero de final")
	titulo = "tablas de multiplicar del " + strconv.Itoa(inicio) + " hasta el " + strconv.Itoa(fin)
	util.Encabezado(titulo, 60)

	generarRangoDeTablas(inicio, fin)
}

func generarTabla(numero int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = %v\n", numero, i, producto(numero, i))
	}
}

func generarRangoDeTablas(inicio, fin int) {

	for i := inicio; i <= fin; i++ {
		fmt.Printf("Tabla de multiplicar del %v\n", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%v x %v = %v\n", i, j, producto(i,j))
		}
	}
}

func producto(a, b int) int {
	return a * b
}
