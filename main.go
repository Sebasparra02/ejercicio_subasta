package main

import "fmt"

func main() {
	var valor_inicial int
	var valor_actual int
	incremento := 0
	var nombre_producto string
	var continuar string

	fmt.Println("ingrese el nombre del producto a subastar")
	_, err := fmt.Scanf("%s", &nombre_producto)
	if err != nil {
		panic(err)
	}
	fmt.Println("ingrese el valor inicial del producto")
	_, err = fmt.Scanf("%d", &valor_inicial)
	if err != nil {
		panic(err)
	}

	fmt.Println("Que inicie la subasta!!!")
	valor_actual = valor_inicial

	for {

		fmt.Println("el valor actual es", valor_actual)
		incremento = valor_actual + 50
		fmt.Println("alguien da mas por ", valor_actual)
		fmt.Println("digite si para continuar y no para finalizar")
		_, err = fmt.Scanf("%s", &continuar)
		if err != nil {
			panic(err)
		} else if continuar == "no" {
			fmt.Println("vendido por :", valor_actual)
		}
		if continuar == "si" {
			fmt.Println("el valor actual es", valor_actual)
			incremento = valor_actual + 50
			fmt.Println("alguien da mas por ", valor_actual)
			fmt.Println("digite si para continuar y no para finalizar")
			_, err = fmt.Scanf("%s", &continuar)
			if err != nil {
				panic(err)
			}
		} else if continuar == "no" {
			fmt.Println("vendido por: ", valor_actual)
			break

		}

	}

}
