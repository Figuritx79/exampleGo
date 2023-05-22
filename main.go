package main

import (
	"fmt"
)

func main() {
	var (
		nombre string
	)

	fmt.Print("Ingrese su nombre:")
	fmt.Scanln(&nombre)

	fmt.Printf("Mi nombre es %s", nombre)
}
