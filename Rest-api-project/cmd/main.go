package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Doctor struct {
	name       string `json : "name"`
	id         int    `json : "id"`
	department string `json : "department"`
	hospital   string `json : "hospital"`
	email      string `json : "email"`
}

var doctor = map[string]Doctor{}

func main() {
	http.HandleFunc("/createdoctor", createDoctor)
	fmt.Println("Server Started")
	log.Fatalf("Starting server failed: %v", http.ListenAndServe(":8000", nil))

}
func createDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	dctr := Doctor{}
	err := json.NewDecoder(r.Body).Decode(&dctr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	doctor[dctr.name] = dctr
	doctor[dctr.department] = dctr
	doctor[dctr.hospital] = dctr
	doctor[dctr.email] = dctr

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dctr)

	return

}
