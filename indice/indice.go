package main

import (
	util "arsistema/utilidades"
	"fmt"
	"math"
	"strings"
)

func main() {

	mostrarIMC()
}

func mostrarIMC() {
	var (
		peso   float64
		altura float64
		imc    float64
	)

	util.Encabezado("Calculadora IMC", 60)

	altura = util.LeerNumerosDecimales("Ingresa la altura (mts)")
	peso = util.LeerNumerosDecimales("Ingresa el peso (kgs)")

	imc = calcularIMC(peso, altura)

	fmt.Printf("El indice de masa corporal es %.2f (%v)\n", imc, strings.ToUpper(identificadorIMC(imc)))
	util.Linea(60)
}

func calcularIMC(peso, altura float64) float64 {

	return peso / (math.Pow(altura, 2))
}


func identificadorIMC(imc float64) string {
	
	if imc < 18.5  {
		return "Bajo peso"
	} else if imc >= 18.5 && imc <= 24.99 { 
		return "Peso Normal"
	} else if imc >= 25 && imc <= 29.99 {
		return "Sobrepeso"
	} else {
		return "Obesidad"
	}
}