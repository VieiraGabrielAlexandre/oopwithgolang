package main

import (
	"fmt"
	"gooop/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaGabriel := contas.ContaPoupanca{}
	contaGabriel.Depositar(100)

	contaJuliana := contas.ContaCorrente{}
	contaJuliana.Depositar(150)

	PagarBoleto(&contaGabriel, 50)
	PagarBoleto(&contaJuliana, 50)

	fmt.Println(contaGabriel.ObterSaldo())
	fmt.Println(contaJuliana.ObterSaldo())
}
