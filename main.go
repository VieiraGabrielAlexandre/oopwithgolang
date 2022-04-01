package main

import (
	"fmt"
	c "gooop/contas"
)

func main() {
	contaGabriel := c.ContaCorrente{Titular: "Gabriel", NumeroAgencia: 123, NumeroConta: 456, Saldo: 1000.00}
	contaDaSilvia := c.ContaCorrente{Titular: "Silvia", Saldo: 300}

	//fmt.Println(contaGabriel.Sacar(800.00))
	fmt.Println(contaGabriel.Saldo)
	status, valor := contaGabriel.Depositar(500.00)
	fmt.Println(status, valor)
	fmt.Println(contaGabriel.Saldo)

	statusTransferencia := contaDaSilvia.Transferir(200, &contaGabriel)
	fmt.Println(statusTransferencia)
	fmt.Println(contaDaSilvia.Saldo)
	fmt.Println(contaGabriel.Saldo)

}
