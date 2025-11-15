package main

import "fmt"

type BalanceRepository interface {
	CreateAccount() int

	GetBalance(accountId int) (int, bool)

	SetBalance(accountId int, amount int)
}

type PlainTransactionService struct {
	balance BalanceRepository
}

func (service PlainTransactionService) CreateAccount() int {
	return service.balance.CreateAccount()
}

func (service PlainTransactionService) Deposit(accountId int, moneyAmount int) error {
	if moneyAmount <= 0 {
		return fmt.Errorf("negative money amount (%v) deposit", moneyAmount)
	}

	if balance, ok := service.balance.GetBalance(accountId); ok {
		service.balance.SetBalance(accountId, balance+moneyAmount)
		return nil
	} else {
		return fmt.Errorf("user %v not found", accountId)
	}
}

func (service PlainTransactionService) Withdraw(accountId int, moneyAmount int) error {
	if moneyAmount <= 0 {
		return fmt.Errorf("negative money amount (%v) withdraw", moneyAmount)
	}

	if balance, ok := service.balance.GetBalance(accountId); ok {
		balance -= moneyAmount
		if balance < 0 {
			return fmt.Errorf("not enough money to wtihdraw %v", moneyAmount)
		}

		service.balance.SetBalance(accountId, balance)
		return nil
	} else {
		return fmt.Errorf("user %v not found", accountId)
	}
}

func (service PlainTransactionService) Transfer(fromId int, toId int, moneyAmount int) error {
	if moneyAmount <= 0 {
		return fmt.Errorf("negative money amount (%v) transfer", moneyAmount)
	}

	fromBalance, ok := service.balance.GetBalance(fromId)
	if !ok {
		return fmt.Errorf("user %v not found", fromId)
	}

	toBalance, ok := service.balance.GetBalance(toId)
	if !ok {
		return fmt.Errorf("user %v not found", toId)
	}

	fromBalance -= moneyAmount
	toBalance += moneyAmount

	if fromBalance < 0 {
		return fmt.Errorf("not enough money to transfer %v", moneyAmount)
	}

	service.balance.SetBalance(fromId, fromBalance)
	service.balance.SetBalance(toId, toBalance)

	return nil
}
