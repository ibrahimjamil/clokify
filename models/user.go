package models

import "database/sql/driver"

type UserType string

const (
	DEVELOPER UserType = "DEVELOPER"
	CLIENT    UserType = "CLIENT"
)

func (ct *UserType) Scan(value interface{}) error {
	*ct = UserType(value.([]byte))
	return nil
}

func (ct UserType) Value() (driver.Value, error) {
	return string(ct), nil
}

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
