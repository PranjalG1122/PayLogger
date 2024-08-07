package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		var user User
		handleError(json.NewDecoder(r.Body).Decode(&user), w)
	}
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/signup", handleSignUp)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
