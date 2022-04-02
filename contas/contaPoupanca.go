package contas

import "gooop/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque < c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Operação realizada com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(depositar float64) (string, float64) {
	if depositar > 0 {
		c.saldo += depositar
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Deposito não realizado", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
