package bankaccount

import "sync"

type Status string

const (
	Unopened Status = "Unopened"
	Active   Status = "Active"
	Closed   Status = "Closed"
)

type Account struct {
	balance int64
	status  string
	mu      sync.Mutex
}

func Open(amt int64) *Account {
	if amt < 0 {
		return nil
	}

	return &Account{
		balance: amt,
		status:  string(Active),
	}
}

func (a *Account) IsActive() bool {
	return a.status == string(Active)
}

func (a *Account) IsInActive() bool {
	return a.status != string(Active)
}

func (a *Account) Balance() (bal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.IsInActive() {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.IsInActive() {
		return 0, false
	}

	if amt < 0 {
		return a.balance, false
	}

	a.balance += amt
	return a.balance, true
}

func (a *Account) Withdraw(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.IsInActive() {
		return 0, false
	}

	if amt < 0 {
		return a.balance, false
	}

	if amt <= a.balance {
		a.balance -= amt
		return a.balance, true
	}
	return a.balance, false
}

func (a *Account) Close() (pay int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.IsInActive() {
		return 0, false
	}

	payout := a.balance
	a.balance = 0
	a.status = string(Closed)
	return payout, true
}
