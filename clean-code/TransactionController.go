package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type TransactionService interface {
	CreateAccount() int
	GetBalance(accountId int) (int, bool)
	GetBalances() map[int]int
	Deposit(accountId int, moneyAmount int) error
	Withdraw(accountId int, moneyAmount int) error
	Transfer(fromId int, toId int, moneyAmount int) error
}

type TransactionController struct {
	service TransactionService
	wg      *sync.WaitGroup
}

type accountData struct {
	Balance int `json:"balance"`
}

func (controller TransactionController) AllAccountsRequests(w http.ResponseWriter, r *http.Request) {
	controller.wg.Add(1)
	defer controller.wg.Done()

	if r.Method == "POST" {
		id := controller.service.CreateAccount()

		data, err := json.Marshal(id)
		if err != nil {
			panic("id should be marshallable")
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	} else if r.Method == "GET" {
		balances := controller.service.GetBalances()

		accounts := map[int]accountData{}
		for id, balance := range balances {
			accounts[id] = accountData{balance}
		}

		data, err := json.Marshal(accounts)
		if err != nil {
			panic("accounts data should be marshallable")
		}

		w.Write(data)
	} else {
		http.Error(w, "only GET and POST method are allowed for this endpoint", http.StatusMethodNotAllowed)
	}
}

func (controller TransactionController) GetAccountData(w http.ResponseWriter, r *http.Request) {
	controller.wg.Add(1)
	defer controller.wg.Done()

	if r.Method != "GET" {
		http.Error(w, "only GET method is allowed for this endpoint", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 1 {
		panic("this request should be handled by HandleAllAccountsRequests")
	}

	id, err := strconv.Atoi(pathParts[len(pathParts)-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	balance, ok := controller.service.GetBalance(id)
	if !ok {
		http.Error(w, fmt.Sprintf("account %v doesn't exist", id), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(accountData{balance})
	if err != nil {
		panic("accountData should be marshallable")
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

type depositRequest struct {
	AccountId   int `json:"account_id"`
	MoneyAmount int `json:"money_amount"`
}

func (controller TransactionController) Deposit(w http.ResponseWriter, r *http.Request) {
	controller.wg.Add(1)
	defer controller.wg.Done()

	if r.Method != "POST" {
		http.Error(w, "only POST method is allowed for this endpoint", http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var parsed depositRequest
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = controller.service.Deposit(parsed.AccountId, parsed.MoneyAmount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type withdrawRequest struct {
	AccountId   int `json:"account_id"`
	MoneyAmount int `json:"money_amount"`
}

func (controller TransactionController) Withdraw(w http.ResponseWriter, r *http.Request) {
	controller.wg.Add(1)
	defer controller.wg.Done()

	if r.Method != "POST" {
		http.Error(w, "only POST method is allowed for this endpoint", http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var parsed withdrawRequest
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = controller.service.Withdraw(parsed.AccountId, parsed.MoneyAmount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type transferRequest struct {
	FromId      int `json:"from_id"`
	ToId        int `json:"to_id"`
	MoneyAmount int `json:"money_amount"`
}

func (controller TransactionController) Transfer(w http.ResponseWriter, r *http.Request) {
	controller.wg.Add(1)
	defer controller.wg.Done()

	if r.Method != "POST" {
		http.Error(w, "only POST method is allowed for this endpoint", http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var parsed transferRequest
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = controller.service.Transfer(parsed.FromId, parsed.ToId, parsed.MoneyAmount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
