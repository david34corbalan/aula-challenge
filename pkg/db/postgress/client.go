package postgress

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(dns string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func CreateDataBase(dbName string, dns string) {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s');", dbName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Base de datos creada exitosamente.")
	} else {
		fmt.Println("La base de datos ya existe.")
	}
}
