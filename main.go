package main

import "fmt"

func main() {
	var valor, aux int
	var num float64
	fmt.Print("Introduce el limite ")
	fmt.Scan(&valor)
	num = 1
	for i := 0; i < valor; i++ {
		aux = 1
		for x := i + 1; x > 1; x-- {
			aux = aux * x
		}
		k:=float64(1)
		k2:=float64(aux)
		num = num + k/k2
	}
	fmt.Printf("E  %.10f\n",num)

}
