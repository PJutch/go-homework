package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type TransactionService interface {
	CreateAccount() int
	Deposit(accountId int, moneyAmount int) error
	Withdraw(accountId int, moneyAmount int) error
	Transfer(fromId int, toId int, moneyAmount int) error
}

type TransactionController struct {
	service TransactionService
}

func (controller TransactionController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only POST method is allowed for this endpoint", http.StatusMethodNotAllowed)
		return
	}

	id := controller.service.CreateAccount()

	data, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

type depositRequest struct {
	accountId   int
	moneyAmount int
}

func (controller TransactionController) Deposit(w http.ResponseWriter, r *http.Request) {
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

	err = controller.service.Deposit(parsed.accountId, parsed.moneyAmount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type withdrawRequest struct {
	accountId   int
	moneyAmount int
}

func (controller TransactionController) Withdraw(w http.ResponseWriter, r *http.Request) {
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

	err = controller.service.Withdraw(parsed.accountId, parsed.moneyAmount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type transferRequest struct {
	fromId, toId int
	moneyAmount  int
}

func (controller TransactionController) Transfer(w http.ResponseWriter, r *http.Request) {
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

	err = controller.service.Transfer(parsed.fromId, parsed.toId, parsed.moneyAmount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
