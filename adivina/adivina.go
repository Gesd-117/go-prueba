package main

import (
	util "arsistema/utilidades"
	"fmt"
	"math/rand/v2"
	"strings"
)

func main() {

	adivinarNumero()

}

func adivinarNumero() {

	var (
		secreto        int
		numero         int
		alias          string
		intentos       int
		numeroIntentos int
		maximoIntentos int //cantidad de intentos -1
	)

	secreto = rand.IntN((10) + 1)

	util.Encabezado("juego de adivina el numero", 60)
	alias = util.LeerTexto("ingrese su alias", "alias")
	alias = strings.ToUpper(alias)
	maximoIntentos = 4

	for numero != secreto {
		numero = util.LeerNumeroEntero("Ingresa un numero")
		intentos++
		numeroIntentos = maximoIntentos - intentos

		if intentos == maximoIntentos {
			break
		}
		if numero < secreto {
			mensajeForIf("el numero es menor que el secreto", alias, numeroIntentos)
		} else if numero > secreto {
			mensajeForIf("el numero es mayor que el secreto", alias, numeroIntentos)
		} else {
			fmt.Printf("Felicidades %v adivinaste el numero 🎉\n", alias)
		}

	}

	fmt.Println("Se te acabaron los intentos")

}

func mensajeForIf(msj string, alias string, intentos int) {

	var mensajes [2]string 
		mensajes[0] = "te queda"
		mensajes[1] = "intento"

	if intentos == 1 {
		fmt.Printf("%v, %v... %v %v %v\n", alias, msj, mensajes[0], intentos, mensajes[1])
	} else {
		fmt.Printf("%v, %v... te quedan %v intentos\n", alias, msj, intentos)
	}
}
