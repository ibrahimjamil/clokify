package models

type User struct {
	ID       int
	Name     string    `db:"name"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Type     string    `db:"type"`
	Projects []Project `gorm:"many2many:project_users;"`
}

func GetUserStruct() *User {
	return &User{}
}
