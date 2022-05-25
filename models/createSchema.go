package models

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

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
