package models

type UserType string

const (
	DEVELOPER UserType = "DEVELOPER"
	CLIENT    UserType = "CLIENT"
)

type User struct {
	ID       int
	Name     string    `db:"name"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Type     UserType  `sql:"type:ENUM('DEVELOPER', 'CLIENT')" gorm:"column:user_type"`
	Projects []Project `gorm:"many2many:project_users;"`
}

func GetUserStruct() *User {
	return &User{}
}
