package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type State struct {
	Counter int `json:"counter"`
}

func main() {

	http.HandleFunc("/counter", handleCounter())
	http.Handle("/", http.FileServer(http.Dir("./build")))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}

func handleCounter() func(w http.ResponseWriter, r *http.Request) {
	state := &State{
		Counter: 5,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(*state)
		if err != nil {
			fmt.Println("unable to marshal state", err)
		}

		switch r.Method {
		case "GET":
			_, err := w.Write(b)
			if err != nil {
				fmt.Println("was not able to write state", err)
			}

		default:
			_, err := w.Write([]byte("Method not implemented"))
			if err != nil {
				fmt.Println("was not able to write state", err)
			}
		}

	}

}
