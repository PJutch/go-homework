package main

type InMemoryBalanceRepository struct {
	nextAccountId int
	balances      map[int]int
}

func (repo *InMemoryBalanceRepository) CreateAccount() int {
	id := repo.nextAccountId
	repo.nextAccountId += 1
	return id
}

func (repo *InMemoryBalanceRepository) GetBalance(accountId int) (int, bool) {
	balance, ok := repo.balances[accountId]
	return balance, ok
}

func (repo *InMemoryBalanceRepository) SetBalance(accountId int, amount int) {
	repo.balances[accountId] = amount
}
