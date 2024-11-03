package db

import (
	"backend/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	return db
}

func InitDB(config config.PgConnectionConfig) {
	var err error
	dsn := "postgres://%s:%s@%s:%s/%s"
	dsn = fmt.Sprintf(
		dsn,
		config.User,
		config.Password,
		config.ContainerName,
		config.Port,
		config.Database,
	)
	fmt.Println("dsn = ", dsn)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	err = db.AutoMigrate(
		&User{},
		&PaymentMethod{},
		&BillPosition{},
		&Bill{},
		&Result{},
		&Group{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
