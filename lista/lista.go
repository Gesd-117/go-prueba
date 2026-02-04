package main

import (
	util "arsistema/utilidades"
	"fmt"
	"slices"
	_ "strings"
)

func main() {
	edades := []int{
		21,
		17,
		15,
		23,
		25,
		16,
		15,
		12,
		13,
		17,
		30,
	}
	edades = append(edades, 30)
	edades = append(edades, 33)

	util.Encabezado("Leer la lista de edades", 50)
	listaDeEdades(edades...)

	listaDeEdadesAsc(edades...)
	listaDeEdadesDes(edades...)
	edadMaxMin(edades...)
}

func listaDeEdades(edades ...int) {
	for indice, edad := range edades {
		fmt.Printf("La edad numero %v es: %v\n", indice+1, edad)
	}
}

func listaDeEdadesAsc(edades ...int) {
	util.Encabezado("lista de edades ordenada ascendentemente", 70)
	slices.Sort(edades)
	listaDeEdades(edades...)
}

func listaDeEdadesDes(edades ...int) {
	util.Encabezado("lista de edades ordenada descendentemente", 70)
	slices.SortFunc(edades, func(a, b int) int {
		return b - a
	})
	listaDeEdades(edades...)
}

func edadMaxMin(edades ...int) {

	util.Encabezado("Edad mas maxima y minima", 70)

	var (
		edadMinima int
		edadMaxima int
	)

	for indice, edad := range edades {

		if indice == 0 {
			edadMaxima = edad
			edadMinima = edad
		}

		if edadMaxima < edad {
			edadMaxima = edad
		}
		if edadMinima > edad {
			edadMinima = edad
		}

	}

	fmt.Printf("La edad maxima  es: %v\n", edadMaxima)
	fmt.Printf("La edad minima es: %v\n", edadMinima)
	util.Linea(70)

}

func listaMayoresDeEdad(edades ...int)  {
	
	util.Encabezado("lista de mayores de edad", 50)
	listaDeEdades(edades...)
	
}