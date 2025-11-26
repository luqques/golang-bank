package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		fmt.Println("Depósito realizado com sucesso")
	}
}

func main() {
	conta := ContaCorrente{
		titular:       "Luques",
		numeroAgencia: 123,
		numeroConta:   456789,
		saldo:         1000.0,
	}

	conta2 := ContaCorrente{"Helena", 321, 987654, 2500.0}

	fmt.Println(conta)
	fmt.Println(conta2)

	var conta3 *ContaCorrente
	conta3 = new(ContaCorrente)
	conta3.titular = "Carlos"
	conta3.numeroAgencia = 111
	conta3.numeroConta = 222333
	conta3.saldo = 500.0

	fmt.Println(*conta3)

	fmt.Println(conta.Sacar(100))
	conta.Depositar(500)
}
