package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	fmt.Println(ContaCorrente{"Gabriel", 123, 456, 1000.00})
}
