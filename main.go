package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Pong struct {
	Status int
	Result string
}

type Rarity string

const (
	RarityN   Rarity = "N"
	RarityR   Rarity = "R"
	RaritySR  Rarity = "SR"
	RaritySSR Rarity = "SSR"
	RarityUR  Rarity = "UR"
	RarityLR  Rarity = "LR"
)

type Item struct {
	Name   string
	Rarity Rarity
}

var card map[string]Item = map[string]Item{
	"ki":   {Name: "木の剣", Rarity: RarityN},
	"isi":  {Name: "石の剣", Rarity: RarityR},
	"tetu": {Name: "鉄の剣", Rarity: RaritySR},
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

func drawall(w http.ResponseWriter, r *http.Request) {
	var cardlist []Item
	for k, v := range card {
		fmt.Printf("key: %s, value: %v\n", k, v)
		cardlist = append(cardlist, v)
	}
	res, err := json.Marshal(cardlist)
	if err != nil {
		fmt.Println(fmt.Errorf("drawall-json.Marshal: %s", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func draw(w http.ResponseWriter, r *http.Request) {
	var cardlist []Item
	for k, v := range card {
		fmt.Printf("key: %s, value: %v\n", k, v)
		cardlist = append(cardlist, v)
	}
	rand.Seed(time.Now().UnixNano())
	ran := rand.Intn(100)
	var ca Item
	switch {
	case ran <= 80:
		ca = cardlist[0]
	case ran <= 95:
		ca = cardlist[1]
	case ran <= 100:
		ca = cardlist[2]
	}
	res, err := json.Marshal(ca)
	if err != nil {
		fmt.Println(fmt.Errorf("draw-json.Marshal: %s", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/", pinghandle)
	http.HandleFunc("/drawall", drawall)
	http.HandleFunc("/draw", draw)

	fmt.Println("server start!")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
