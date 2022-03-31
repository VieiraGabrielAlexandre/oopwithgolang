package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaGabriel := ContaCorrente{"Gabriel", 123, 456, 1000.00}
	contaTeste := ContaCorrente{titular: "Teste", saldo: 150.0}

	fmt.Println(contaGabriel)
	fmt.Println(contaTeste)
}
