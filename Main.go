package main

import (
	"encoding/json"
	"net/http"
)

type Kelompok3 struct {
	Nama string `json:"nama"`
	NIM  string `json:"nim"`
}

func main() {
	kelompok := []Kelompok3{
		{"Alvian Akbar Aulia", "1304212113"},
		{"Muhammad Saladin Fikri Abdulloh", "1304212121"},
		{"Muhammad Sirojul Fu’ad", "1304212094"},
		{"Syahratul Muthi’ah M. Masiming", "1304211013"},
	}

	http.HandleFunc("/kelompok", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(kelompok)
	})

	http.ListenAndServe(":8080", nil)
}
