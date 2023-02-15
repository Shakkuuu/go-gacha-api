package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Pong struct {
	Status int
	Result string
}

func pinghandle(w http.ResponseWriter, r *http.Request) {
	pong := Pong{http.StatusOK, "ok"}

	resp, err := json.Marshal(pong)
	if err != nil {
		fmt.Println(fmt.Errorf("json.Marshal: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func main() {
	http.HandleFunc("/", pinghandle)

	fmt.Println("server start!")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
