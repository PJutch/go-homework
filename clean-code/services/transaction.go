package services

import "fmt"

type BalanceRepository interface {
	CreateAccount() int
	GetBalance(accountId int) (int, bool)
	GetBalances() map[int]int
	SetBalance(accountId int, amount int)
}

type PlainTransactionService struct {
	Repo BalanceRepository
}

func (service PlainTransactionService) CreateAccount() int {
	return service.Repo.CreateAccount()
}

func (service PlainTransactionService) GetBalance(accountId int) (int, bool) {
	return service.Repo.GetBalance(accountId)
}

func (service PlainTransactionService) GetBalances() map[int]int {
	return service.Repo.GetBalances()
}

func (service PlainTransactionService) Deposit(accountId int, moneyAmount int) error {
	if moneyAmount <= 0 {
		return fmt.Errorf("negative money amount (%v) deposit", moneyAmount)
	}

	if balance, ok := service.Repo.GetBalance(accountId); ok {
		service.Repo.SetBalance(accountId, balance+moneyAmount)
		return nil
	} else {
		return fmt.Errorf("user %v not found", accountId)
	}
}

func (service PlainTransactionService) Withdraw(accountId int, moneyAmount int) error {
	if moneyAmount <= 0 {
		return fmt.Errorf("negative money amount (%v) withdraw", moneyAmount)
	}

	if balance, ok := service.Repo.GetBalance(accountId); ok {
		balance -= moneyAmount
		if balance < 0 {
			return fmt.Errorf("not enough money to wtihdraw %v", moneyAmount)
		}

		service.Repo.SetBalance(accountId, balance)
		return nil
	} else {
		return fmt.Errorf("user %v not found", accountId)
	}
}

func (service PlainTransactionService) Transfer(fromId int, toId int, moneyAmount int) error {
	if moneyAmount <= 0 {
		return fmt.Errorf("negative money amount (%v) transfer", moneyAmount)
	}

	fromBalance, ok := service.Repo.GetBalance(fromId)
	if !ok {
		return fmt.Errorf("user %v not found", fromId)
	}

	toBalance, ok := service.Repo.GetBalance(toId)
	if !ok {
		return fmt.Errorf("user %v not found", toId)
	}

	fromBalance -= moneyAmount
	toBalance += moneyAmount

	if fromBalance < 0 {
		return fmt.Errorf("not enough money to transfer %v", moneyAmount)
	}

	service.Repo.SetBalance(fromId, fromBalance)
	service.Repo.SetBalance(toId, toBalance)

	return nil
}
