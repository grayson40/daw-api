package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grayson40/daw-api/db"
	"github.com/grayson40/daw-api/types"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetUsers(w, r)
	case http.MethodPost:
		handleAddUser(w, r)
	case http.MethodPut:
		// Handle PUT request for users
	case http.MethodDelete:
		// Handle DELETE request for users
	default:
		// Handle other HTTP methods or return an error
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

// Get a list of all users from the database
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// Add a new user to the database
func handleAddUser(w http.ResponseWriter, r *http.Request) {
	// Parse the new user from the request body
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the new user to the database
	createdID, err := db.AddUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the ID of the newly created user
	json.NewEncoder(w).Encode(createdID)
}
