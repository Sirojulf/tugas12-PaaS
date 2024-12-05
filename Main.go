package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Kelompok3 struct {
	Nama string `json:"nama"`
	NIM  string `json:"nim"`
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Acces-Control-Allow-Origin", "*")
		w.Header().Set("Acces-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Acces-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	kelompok := []Kelompok3{
		{"Alvian Akbar Aulia", "1304212113"},
		{"Muhammad Saladin Fikri Abdulloh", "1304212121"},
		{"Muhammad Sirojul Fu’ad", "1304212094"},
		{"Syahratul Muthi’ah M. Masiming", "1304211013"},
	}

	kelompokHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(kelompok)
	})

	mux := http.NewServeMux()
	mux.Handle("/kelompok", loggingMiddleware(corsMiddleware(kelompokHandler)))

	log.Printf("Server running port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start :8080 %v\n", err)
	}
}
