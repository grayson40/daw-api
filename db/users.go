package db

import (
	"github.com/grayson40/daw-api/types"
	_ "github.com/mattn/go-sqlite3"
)

// Adds a new user to the database
func AddUser(newUser types.User) (int, error) {
	// Prepare the SQL statement
	stmt, err := DB.Prepare("INSERT INTO users (username, email) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the statement
	res, err := stmt.Exec(newUser.Username, newUser.Email)
	if err != nil {
		return 0, err
	}

	// Get the ID of the newly inserted user
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Returns a list of all users from the database
func GetUsers() ([]types.User, error) {
	// Prepare the SQL statement
	stmt, err := DB.Prepare("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the statement
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and add each user to the list
	var users []types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
