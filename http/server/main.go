package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func randIntInRange(from int /* inclusive */, to int /* exclusive */) int {
	return rand.Intn(to-from) + from
}

func randFloatInRange(from float64 /* inclusive */, to float64 /* exclusive */) float64 {
	return rand.Float64()*(from-to) + to
}

func randBool(trueChance float64) bool {
	return rand.Float64() < trueChance
}

type ServerState struct {
	server http.Server
	wg     sync.WaitGroup
}

func (state *ServerState) handleVersionRequest(w http.ResponseWriter, r *http.Request) {
	state.wg.Add(1)
	defer state.wg.Done()

	if r.Method == http.MethodGet {
		fmt.Fprint(w, "0.0.1")
	}
}

type jsonInput struct {
	InputString string `json:"inputString"`
}

type jsonOutput struct {
	OutputString string `json:"outputString"`
}

func (state *ServerState) handleDecodeRequest(w http.ResponseWriter, r *http.Request) {
	state.wg.Add(1)
	defer state.wg.Done()

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var input jsonInput
		err2 := json.Unmarshal(body, &input)
		if err2 != nil {
			fmt.Println(err2)
			return
		}

		outputBytes, err := base64.StdEncoding.DecodeString(input.InputString)
		if err != nil {
			fmt.Println(err)
			return
		}

		response, err := json.Marshal(jsonOutput{string(outputBytes)})
		if err != nil {
			fmt.Println(err)
			return
		}

		w.Write(response)
	}
}

func (state *ServerState) handleHardOp(w http.ResponseWriter, r *http.Request) {
	state.wg.Add(1)
	defer state.wg.Done()

	if r.Method == http.MethodGet {
		time.Sleep(time.Duration(randFloatInRange(10, 20)) * time.Second)

		if randBool(0.5) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(randIntInRange(500, 527))
		}
	}
}

func main() {
	fmt.Println("Starting...")

	state := ServerState{http.Server{Addr: ":8080"}, sync.WaitGroup{}}

	http.HandleFunc("/version", state.handleVersionRequest)
	http.HandleFunc("/decode", state.handleDecodeRequest)
	http.HandleFunc("/hard-op", state.handleHardOp)

	go func() {
		fmt.Println("Started")

		err := state.server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := state.server.Shutdown(ctx); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Server shut down")

	state.wg.Wait()
	fmt.Println("Wg waited out")
}
