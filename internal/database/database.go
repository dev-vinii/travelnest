// Preciso fazer a conexão usando postres

package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "example"
	password = "example"
	dbname   = "example"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil
}