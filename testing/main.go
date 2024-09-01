package main

import (
	"crypto/sha1"
	"fmt"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

)

var urlStore = struct {
	sync.RWMutex
	mappings map[string] string 
}{mappings: make(map[string]string)}

func generateshortURL (longURL string) string {
	h := sha1.New()
	h.Write([]byte(longURL))
	return hex.EncodeToString(h.Sum(nil))[:8]
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		LongURL string `json:"long_url"`
	} 
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if request.LongURL == "" {
		http.Error(w, "Long URL is needed", http.StatusBadRequest)
		return
	}
	shortURL := generateshortURL(request.LongURL)
	urlStore.Lock()
	urlStore.mappings[shortURL] = request.LongURL
	urlStore.Unlock()
	response := struct {
		
	} 
}

func main() {

stringtemplate := "Hello, my name is: "   
name := "Ankit"  
out := fmt.Sprintf("%s%S", stringtemplate , name) // now we can use out var 
}
