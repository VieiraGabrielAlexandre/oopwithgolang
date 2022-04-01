package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque < c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(depositar float64) (string, float64) {
	if depositar > 0 {
		c.Saldo += depositar
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Deposito não realizado", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor > 0 && valor < c.Saldo {
		c.Saldo -= valor
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}
