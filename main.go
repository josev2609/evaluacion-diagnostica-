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

	for opcion = 3 {

	}