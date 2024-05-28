package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"todoList/ent"

	_ "github.com/go-sql-driver/mysql"
)

var DB *ent.Client

func InitDB() error {
	db, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	err = db.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("filed creating schema resources: %v", err)
	}

	DB = db
	return nil
}
