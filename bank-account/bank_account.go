package account

import (
	"sync"
)

type Account struct {
	amount int64
	open bool
	mutex sync.Mutex
}

func Open(initialAmount int64) *Account {
	if initialAmount < 0 {
		return nil
	}
	return &Account{
		amount: initialAmount,
		open: true,
		mutex: sync.Mutex{},
	}
}

func (account *Account) Balance() (int64, bool) {
	account.mutex.Lock()
	defer account.mutex.Unlock()
	if !account.open {
		return 0, false
	}
	return account.amount, account.open
}

func (account *Account) Close() (int64, bool) {
	account.mutex.Lock()
	defer account.mutex.Unlock()
	if !account.open {
		return 0, false
	}
	amount := account.amount
	account.amount = 0
	account.open = false
	return amount, true
}

func (account *Account) Deposit(amount int64) (int64, bool) {
	account.mutex.Lock()
	defer account.mutex.Unlock()
	if !account.open || amount + account.amount < 0 {
		return 0, false
	}
	account.amount += amount
	return account.amount, true
}