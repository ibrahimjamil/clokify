package db

import (
	. "clokify/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() (*gorm.DB, error) {
	var config = EnvConfig()
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
