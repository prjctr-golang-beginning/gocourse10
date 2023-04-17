package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{
	{ID: "1", Username: "john_doe", Email: "john.doe@example.com"},
	{ID: "2", Username: "jane_doe", Email: "jane.doe@example.com"},
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

	// exercises
	// 1. Створіть API: отримати користувача по ID; Змінити статус існуючого користувача (active, banned). Дані про користувачів зберігайте у пам'яті
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.NotFound(w, r)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)

			var updatedUser User
			err := json.NewDecoder(r.Body).Decode(&updatedUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updatedUser.ID = params["id"]
			users = append(users, updatedUser)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}
