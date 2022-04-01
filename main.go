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

	//fmt.Println(contaGabriel.Sacar(800.00))
	fmt.Println(contaGabriel.saldo)
	status, valor := contaGabriel.Depositar(500.00)
	fmt.Println(status, valor)
	fmt.Println(contaGabriel.saldo)
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque < c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(depositar float64) (string, float64) {
	if depositar > 0 {
		c.saldo += depositar
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Deposito não realizado", c.saldo
	}
}
