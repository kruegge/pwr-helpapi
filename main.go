package main

import (
	"fmt"
	"helpapi/pkg/malo"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
)

func handleMalo(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(malo.Create()))
}

func handleMain(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("powerAPI"))
}

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func router() {
	http.HandleFunc("/malo", handleMalo)
	http.HandleFunc("/", handleMain)
}

func main() {
	router()
	if inLambda() {
		fmt.Println("running aws lambda")
		log.Fatal(gateway.ListenAndServe("", nil))
	} else {
		fmt.Println("running local")
		log.Fatal(http.ListenAndServe(":8082", nil))
	}
}
