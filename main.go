package main

import (
	"fmt"

	contas "github.com/marianavif/go-oo-alura/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)

	fmt.Println(status)

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
