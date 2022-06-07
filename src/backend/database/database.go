package database

import (
	"log"
	"os"

	"github.com/NQuirmbach/todo-app/entity"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.Todo{})

	return err
}

func Init() (db *gorm.DB) {
	dns := os.Getenv("POSTGRES")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Cannot open database")
	}

	err = migrate(db)
	if err != nil {
		panic("Error while migrating database")
	}

	log.Println("Migrated database")
	return
}
