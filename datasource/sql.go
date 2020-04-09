package datasource

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"path/filepath"
)

var Db *sqlx.DB

const (
	DB_HOST     = ""
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_NAME     = ""
)

func MustNewDB() {
	wd, _ := os.Getwd()
	sqlPath := filepath.Join(filepath.Dir(wd), "scripts", "sql", "users.sql")
	var err error
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	Db, err = sqlx.Connect("pgx", dbinfo)
	if err != nil {
		log.Fatalln(err)
	}
	_, sqlErr := sqlx.LoadFile(Db, sqlPath)
	if sqlErr != nil {
		log.Fatalf("Can't apply DB schema")
	}
}
