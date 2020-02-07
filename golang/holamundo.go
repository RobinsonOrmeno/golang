package main

import "fmt"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	//gorras()

	fmt.Println(gorrasClosures(45, "$"))
}

func gorras() {

	var gorra_negra = Gorra{"Zico", "Azul", 15.99, true}

	fmt.Println(gorra_negra)

}

func gorrasClosures(pedido float32, moneda string) (string, float32, string) {
	precio := func() float32 {
		return pedido * 7
	}
	return "El precio de gorras Pedidas es:", precio(), moneda

}
