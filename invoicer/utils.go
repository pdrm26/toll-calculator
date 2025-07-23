package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func writeJSON(w http.ResponseWriter, status int, res any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}

func makeStore() Storer {
	storeType := os.Getenv("AGG_STORE_TYPE")
	switch storeType {
	case "memory":
		return NewMemoryStore()
	default:
		log.Fatalf("invalid memory type: %s", storeType)
		return nil
	}
}
