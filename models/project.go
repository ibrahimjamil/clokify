package models

type Project struct {
	ID       int
	Name     string `db:"name"`
	IsPublic string `db:"is_public"`
	ColorTag string `db:"color_tag"`
	UserID   string `db:"user_id"`
	User     User   `gorm:"foreignKey:user_id"`
}

func GetProjectStruct() *Project {
	return &Project{}
}
