package main

import (
	"fmt"
	util "arsistema/utilidades"
	_"strings"
	_"slices"
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
	
	util.Linea(50)
	fmt.Println("la cantidad de edades es", len(edades))
	util.Linea(50)
	fmt.Println("la tercera edad es", edades[2])
	util.Linea(50)
	fmt.Println("la quinta edad es", edades[4])
	util.Linea(50)
	fmt.Println("la ultima edad es", edades[len(edades)-1])
	util.Linea(50)

	edades = append(edades, 30)
	edades = append(edades, 33)
	
	fmt.Println("la ultima edad es", edades[len(edades)-1])
	util.Linea(50)
}