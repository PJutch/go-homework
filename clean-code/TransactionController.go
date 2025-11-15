package main

type TransactionService interface {
	CreateAccount() int
	Deposit(accountId int, moneyAmount int)
	Withdraw(accountId int, moneyAmount int)
	Transfer(fromId int, toId int, moneyAmount int)
}
