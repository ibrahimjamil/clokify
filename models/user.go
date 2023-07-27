package models

type User struct {
	ID       int
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	IsActive bool   `db:"is_active"`
	Type     string `db:"type"`
}

func GetUserStruct() *User {
	return &User{}
}
