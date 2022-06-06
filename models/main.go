package models

import (
	"context"
	"os"

	"github.com/NoamBoni/gofoloapp/helpers"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Roles struct {
	T string
	P string
}

var Role Roles
var Db *pg.DB

func init() {
	Db = ConnectDB()
	Role = Roles{
		T: "Therapist",
		P: "Patient",
	}
}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Patient)(nil),
		(*Meeting)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func ConnectDB() *pg.DB {
	helpers.LoadEnv()
	db := pg.Connect(&pg.Options{
		Addr:      os.Getenv("DB_URL"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASS"),
		Database:  os.Getenv("DB_NAME"),
		TLSConfig: nil,
	})
	ctx := context.Background()
	if _, err := db.ExecContext(ctx, "SELECT 1"); err != nil {
		panic(err)
	}
	CreateSchema(db)
	return db
}
