package sql

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewSql() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:example@/go_playground")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
