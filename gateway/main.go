package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "the listen address of the http server")
	flag.Parse()

	http.HandleFunc("/invoice", makeAPIFunc(handleGetInvoice))

	logrus.Infof("gateway HTTP server is up and running on port %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))

}

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	return writeJSON(w, http.StatusOK, map[string]string{"data": "ok"})
}

func writeJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}
