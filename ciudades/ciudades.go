package main

import (
	util "arsistema/utilidades"
	"fmt"
	"slices"
	"strings"
)

func main()  {
		
	ciudades := []string{}
	avisos := [6]string{
		"primera",
		"segunda",
		"tercera",
		"cuarta",
		"quinta",
		"sexta",
	}

	var ciudad string
	const nroDeCiudades int = 6

	util.Encabezado("ingreso de ciudades",46)
	for i := range nroDeCiudades {
		ciudad = util.LeerTexto("Ingrese la " + avisos[i] + " ciudad", "ciudad")
		ciudades = append(ciudades, ciudad)
	}

	util.Encabezado("lista de ciudades",50)
	for indice, ciudad := range ciudades {		
		
		fmt.Printf("La %v ciudad es %v\n", avisos[indice], strings.ToUpper(ciudad))
	}

	util.Encabezado("lista de ciudades ordenada ascendentemente",50)
	slices.Sort(ciudades)
	for indice, ciudad := range ciudades {		
		
		fmt.Printf("La %v ciudad es %v\n", avisos[indice], strings.ToUpper(ciudad))
	}

	util.Encabezado("lista de ciudades ordenada descendentemente",50)
	slices.SortFunc(ciudades, func(a,b string) int {
		return strings.Compare(b, a)
	})
	for indice, ciudad := range ciudades {		
		
		fmt.Printf("La %v ciudad es %v\n", avisos[indice], strings.ToUpper(ciudad))
	}
}
