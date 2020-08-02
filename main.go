package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"housing-anywhere/models"
	"housing-anywhere/services"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func navigation(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var location models.Location
	err = json.Unmarshal(body, &location)
	if err != nil{
		http.Error(w, "Error parsing body content", http.StatusBadRequest)
		return
	}

	loc := services.Calculate(location)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(loc)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", navigation).Methods(http.MethodPost)
	api.HandleFunc("/health", healthcheck).Methods(http.MethodGet)

	log.Print("Listening on localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
