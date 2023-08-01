package models

type ProjectCreateType struct {
	ID       int
	Name     string `db:"name"`
	IsPublic bool   `db:"is_public"`
	ColorTag string `db:"color_tag"`
	UserId   int
}
type Project struct {
	ID       int
	Name     string `db:"name"`
	IsPublic bool   `db:"is_public"`
	ColorTag string `db:"color_tag"`
	Users    []User `gorm:"many2many:project_users;"` // Intermediate table name: project_users
}

func GetProjectStruct() *Project {
	return &Project{}
}
