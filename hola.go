package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

const (
	Pi = 3.14 // Constante de ejemplo
)
const (
	Domingo   = iota // 0
	Lunes            // 1
	Martes           // 2
	Miercoles        // 3
	Jueves           // 4
	Viernes          // 5
	Sabado           // 6
)

func main() {
	fmt.Println("Hola, mundo!")
	fmt.Println(quote.Hello())
	//declaracion de variables
	var firstName, lastName string = "John", "Doe"
	var age int = 10
	var isActive bool = true
	var height float64 = 1.75
	var valorCadena string = "Hola, mundo!"
	fmt.Printf("Nombre: %s %s, Edad: %d\n y valor de pi %f", firstName, lastName, age, Pi)
	fmt.Printf("\nDía de la semana: %d", Lunes) // Imprime el valor de Lunes
	fmt.Printf("\nValor de la cadena: %s", valorCadena)
	fmt.Printf("\nValor del booleano: %t", isActive) // Imprime el valor de Pi
	fmt.Printf("\nValor de la flotante: %f", height) // Imprime el valor de Pi
	s := "100s"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("\nError al convertir la cadena s: %v", err)
	} else {
		fmt.Printf("\nValor de la cadena s convertida a int: %d", i)
	}
}
