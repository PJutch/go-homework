package main

import (
	"clean-code/controllers"
	"clean-code/repositories"
	"clean-code/services"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	server := http.Server{Addr: ":8080"}
	wg := sync.WaitGroup{}

	balanceRepository := repositories.MakeInMemoryBalanceRepository()

	controller := controllers.TransactionController{
		Service: services.PlainTransactionService{
			Repo: &balanceRepository,
		},
		Wg: &wg,
	}

	http.HandleFunc("/accounts", controller.AllAccountsRequests)
	http.HandleFunc("/accounts/", controller.GetAccountData)
	http.HandleFunc("/deposit", controller.Deposit)
	http.HandleFunc("/withdraw", controller.Withdraw)
	http.HandleFunc("/transfer", controller.Transfer)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println(err)
		return
	}

	wg.Wait()
}
