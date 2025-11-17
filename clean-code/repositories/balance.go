package repositories

type InMemoryBalanceRepository struct {
	nextAccountId int
	balances      map[int]int
}

func MakeInMemoryBalanceRepository() InMemoryBalanceRepository {
	return InMemoryBalanceRepository{nextAccountId: 0, balances: map[int]int{}}
}

func (repo *InMemoryBalanceRepository) CreateAccount() int {
	id := repo.nextAccountId
	repo.nextAccountId += 1

	repo.balances[id] = 0
	return id
}

func (repo *InMemoryBalanceRepository) GetBalance(accountId int) (int, bool) {
	balance, ok := repo.balances[accountId]
	return balance, ok
}

func (repo *InMemoryBalanceRepository) SetBalance(accountId int, amount int) {
	repo.balances[accountId] = amount
}

func (repo *InMemoryBalanceRepository) GetBalances() map[int]int {
	return repo.balances
}
