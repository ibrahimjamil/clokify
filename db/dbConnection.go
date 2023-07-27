package db

import (
	. "clokify/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(config *Config) (*gorm.DB, error) {
	dbEnvs := "host=" + config.Host +
		" port=" + config.Port +
		" user=" + config.User +
		" password=" + config.Password +
		" dbname=" + config.DBName

	db, err := gorm.Open(postgres.Open(dbEnvs), &gorm.Config{})
	if err == nil {
		MigrateAllDB(db)
		return db, nil
	}
	return db, err
}
