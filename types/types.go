package types

type User struct {
	ID       int
	Username string
	Email    string
}

type Project struct {
	ID          int
	Name        string
	Description string
}

type Issue struct {
	ID          int
	Title       string
	Description string
	ProjectID   int
	AssigneeID  int
}
