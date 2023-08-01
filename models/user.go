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

type UserRegistrationType struct {
	ID       int      `json:"id" binding:"required"`
	Name     string   `json:"name" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Type     UserType `sql:"type:ENUM('DEVELOPER', 'CLIENT')" gorm:"column:user_type"`
}
