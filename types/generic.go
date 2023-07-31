package types

import "gorm.io/gorm"

type ServiceManager struct {
	Db *gorm.DB
}
