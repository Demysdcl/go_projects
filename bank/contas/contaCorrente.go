package contas

import "bank/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	NumeroAgencia int
	NumeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valor float64) string  {
    podeSacar := valor > 0 && valor <= c.saldo
    if podeSacar {
        c.saldo -= valor
        return "Saque realizado com sucesso"
    } else {
        return "saldo insuficiente"
    }
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64)  {

    if valor > 0 {
        c.saldo += valor
        return "Deposito realizado com sucesso", c.saldo
    } else {
        return "Valor inválido", c.saldo
    }
}

func (c *ContaCorrente) Transferir(valor float64, destino *ContaCorrente) bool  {
    podeTransferir := valor > 0 && valor <= c.saldo
    if(podeTransferir) {
        c.saldo -= valor
        destino.Depositar(valor)
    }
    return podeTransferir
}


func (c *ContaCorrente) GetSaldo() float64  {
    return c.saldo
}
