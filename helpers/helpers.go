package helpers

import (
	"context"
	"log"
	"os"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
)


func LoadEnv()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() *pg.DB {
	LoadEnv()
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_URL"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		TLSConfig: nil,
	})
	ctx := context.Background()
	if _, err := db.ExecContext(ctx, "SELECT 1"); err != nil {
		panic(err)
	}
	models.CreateSchema(db)
	return db
}