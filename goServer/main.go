package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("post /users", createUser)

	fmt.Println("server starting...")
	http.ListenAndServe(":8080", mux)

}

func handleRoot(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Name == "" {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
