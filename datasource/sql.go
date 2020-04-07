package datasource

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB

const (
	DB_HOST     = ""
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_NAME     = ""
)

var schema = `
CREATE TABLE IF NOT EXISTS public.api_users (
uuid varchar not null primary key,
name varchar,
lastname varchar,
birthdate timestamp
);
`

func MustNewDB() {
	var err error
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	Db, err = sqlx.Connect("pgx", dbinfo)
	if err != nil {
		log.Fatalln(err)
	}
	Db.MustExec(schema)
}
