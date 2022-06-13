package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type State struct {
	Counter int `json:"counter"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func main() {
	http.HandleFunc("/state", handleState())
	http.Handle("/", http.FileServer(http.Dir("./public")))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}

func handleState() func(w http.ResponseWriter, r *http.Request) {
	state := &State{
		Counter: 0,
	}
	errState := &Error{}

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
		case "PUT":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = json.Unmarshal(body, state)
			if err != nil {
				errState.Code = 400
				errState.Message = "Input format must be type int"
				errorMessage, _ := json.Marshal(*errState)
				w.WriteHeader(400)
				_, err := w.Write(errorMessage)
				if err != nil {
					return
				}
				return
			}

			back, err := json.Marshal(*state)
			if err != nil {
				errState.Message = "Unable to marshall state struct"
				fmt.Println("Unable to marshall state struct")
				return
			}
			_, err = w.Write(back)

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
