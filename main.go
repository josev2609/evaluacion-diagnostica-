package main

import "fmt"

var productosVendidos []string
var subtotales []float64

func Regisventa(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)
}

func MostrarEstadisticas() {
	if len(subtotales) == 0 {
	}

	total := 0.0

	for _, subtotal := range subtotales {
		total += subtotal
	}

	fmt.Println("Total recaudado", total)
}

func main() {

	productos := []string{"Arroz", "Leche", "Pan"}
	precios := []float64{1.25, 0.95, 0.50}

	var opcion int

	for opcion != 3 {

		fmt.Println("1. Registrar venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Opción: ")
		fmt.Scan(&opcion)

		if opcion == 1 {

			fmt.Println("Productos:")
			for i := 0; i < 3; i++ {
				fmt.Println(i+1, productos[i], "$", precios[i])
			}

			var numero int
			var cantidad int

			fmt.Print("Seleccione producto: ")
			fmt.Scan(&numero)

			fmt.Print("Cantidad: ")
			fmt.Scan(&cantidad)

			RegistrarVenta(
				productos[numero-1],
				precios[numero-1],
				cantidad,
			)

			fmt.Println("Venta registrada.")

		} else if opcion == 2 {

			MostrarEstadisticas()

		} else if opcion == 3 {

			fmt.Println("Programa terminado.")

		} else {

			fmt.Println("Opción incorrecta.")
		}
	}
}