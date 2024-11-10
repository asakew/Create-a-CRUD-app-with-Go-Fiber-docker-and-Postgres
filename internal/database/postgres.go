package database

import (
	"goFiber/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var FactDB *gorm.DB

func ConnectFDB() {
	var err error
	dsn := "host=localhost user=postgres password=Root dbname=facts_db port=5432 sslmode=disable"
	FactDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Migrate()

}

func Migrate() {
	err := FactDB.AutoMigrate(&models.Fact{})
	if err != nil {
		return
	}
}
