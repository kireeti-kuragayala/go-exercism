//go:build ignore

package main

import (
	"bankaccount"
	"fmt"
)

func main() {
	acc := bankaccount.Open(100)

	bal, ok := acc.Balance()
	fmt.Printf("Balance: %v, %v\n", bal, ok)

	bal, ok = acc.Deposit(150)
	fmt.Printf("Deposited 150: %v, %v\n", bal, ok)

	bal, ok = acc.Withdraw(111)
	fmt.Printf("Withdraw within limit: %v, %v\n", bal, ok)

	bal, ok = acc.Withdraw(761)
	fmt.Printf("Withdraw more than balance: %v, %v\n", bal, ok)

	payout, ok := acc.Close()
	fmt.Printf("Account is closed, payout: %v, %v\n", payout, ok)

	bal, ok = acc.Deposit(120)
	fmt.Printf("Attempt deposit in a closed account: %v, %v\n", bal, ok)
}
