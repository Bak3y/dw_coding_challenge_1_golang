package main

import "fmt"

type balance struct {
	BalanceType string
	Balance     float32
}

func CFIM_resolve_balance(debit_balance, credit_balance float32) balance {
	debit := debit_balance - credit_balance
	credit := credit_balance - debit_balance

	b := &balance{}

	if debit == 0 {
		b.BalanceType = "None"
		b.Balance = 0
	} else if debit < 0 {
		b.BalanceType = "Credit"
		b.Balance = credit
	} else {
		b.BalanceType = "Debit"
		b.Balance = debit
	}

	return *b
}

func main() {
	bal := CFIM_resolve_balance(125.22, 65.08)
	fmt.Print(bal)
}
