package main

import (
	"bank/clientes"
	"bank/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
    conta.Sacar(valorBoleto)
}

type verificarConta interface {
    Sacar(valor float64) string
}


func main() {
    clienteDemys := clientes.Titular{
        Nome: "Demys",
        CPF: "12458932524",
        Profissao: "Dev",
    }

	contaDemys := contas.ContaCorrente{
		Titular: clienteDemys,
		NumeroAgencia: 0001,
		NumeroConta: 00201324,
	}

    contaDemys.Depositar(10000)

    contaPoupancaDemys := contas.ContaPoupanca{
		Titular: clienteDemys,
		NumeroAgencia: 0001,
		NumeroConta: 00201324,
        Operacao: 1515,
	}

    contaPoupancaDemys.Depositar(8000)

    clienteVirginia := clientes.Titular{"Virginia", "2578445415", "Dev"}

    contaVirginia := contas.ContaPoupanca{
		Titular: clienteVirginia,
		NumeroAgencia: 0001,
		NumeroConta: 00201324,
        Operacao: 12450,
	}

    contaVirginia.Depositar(2000)

    contaVirginia.Transferir(100, &contaPoupancaDemys)

    contaDemys.Sacar(100)
    contaVirginia.Depositar(200)

    PagarBoleto(&contaDemys, 20)
    PagarBoleto(&contaPoupancaDemys, 20)
    PagarBoleto(&contaVirginia, 20)

	fmt.Println("contaDemys", contaDemys.GetSaldo())
	fmt.Println("contaPoupancaDemys", contaPoupancaDemys.GetSaldo())
	fmt.Println("contaVirginia", contaVirginia.GetSaldo())



}
