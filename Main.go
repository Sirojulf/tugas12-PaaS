package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Kelompok3 struct {
	Nama string `json:"nama"`
	NIM  string `json:"nim"`
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[START] %s %s %s", r.RemoteAddr, r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("[END] %s %s %s - Duration: %v", r.RemoteAddr, r.Method, r.URL.Path, duration)
	})
}

func main() {
	kelompok := []Kelompok3{
		{"Alvian Akbar Aulia", "1304212113"},
		{"Muhammad Saladin Fikri Abdulloh", "1304212121"},
		{"Muhammad Sirojul Fu’ad", "1304212094"},
		{"Syahratul Muthi’ah M. Masiming", "1304211013"},
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	kelompokHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(kelompok); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	})

	healthHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux := http.NewServeMux()
	mux.Handle("/kelompok", loggingMiddleware(corsMiddleware(kelompokHandler)))
	mux.Handle("/health", healthHandler)

	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Could not start server on port %s: %v\n", port, err)
	}
}
