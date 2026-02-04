package main

import (
	util "arsistema/utilidades"
	"fmt"
	"slices"
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
	listaMayoresDeEdad(edades...)
	listaMenoresDeEdad(edades...)
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

	util.Encabezado("Edad maxima y minima", 50)

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

	fmt.Printf("La edad maxima es: %v\n", edadMaxima)
	fmt.Printf("La edad minima es: %v\n", edadMinima)
	util.Linea(50)

}

func listaMayoresDeEdad(edades ...int) {

	var (
		mayoresDeEdad = []int{}
	)

	for _, edad := range edades {

		if edad >= 18 {
			mayoresDeEdad = append(mayoresDeEdad, edad)
		}

	}

	util.Encabezado("lista de mayores de edad", 50)
	listaDeEdades(mayoresDeEdad...)
	util.Linea(50)

}

func listaMenoresDeEdad(edades ...int) {
	var (
		menoresDeEdad = []int{}
	)

	for _, edad := range edades {

		if edad < 18 {
			menoresDeEdad = append(menoresDeEdad, edad)
		}

	}

	util.Encabezado("lista de menores de edad", 50)
	listaDeEdades(menoresDeEdad...)
	util.Linea(50)
}
