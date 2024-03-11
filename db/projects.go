package db

import "github.com/grayson40/daw-api/types"

// Adds a new project to the database
func AddProject(newProject types.Project) (int, error) {
	// Prepare the SQL statement
	stmt, err := DB.Prepare("INSERT INTO projects (name, description) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the statement
	res, err := stmt.Exec(newProject.Name, newProject.Description)
	if err != nil {
		return 0, err
	}

	// Get the ID of the newly inserted project
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
