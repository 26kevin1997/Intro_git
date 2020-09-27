package main

import "fmt"

func main()  {
	var a float64
	var b float64
	fmt.Println("Cuanto mide un lado del cuadrado: ")
	fmt.Scanf("%f", &a)
	fmt.Println("El area es: ",a*a)
	fmt.Println("Fin del programa")
	fmt.Println("Otra modificacion")
	fmt.Println("Que edad tienes")
	fmt.Scan(&b)
	fmt.Println("La edad es: ", b)
}