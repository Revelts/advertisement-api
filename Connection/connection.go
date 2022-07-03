package Connection

//setup connection
//rabbit mq
//redis
//postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "Revelt"
	password = ""
	dbname   = "trainingapi"
)

var PostgresConnection *sql.DB

func postgresConn() (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	return
}

func InitializeConnection() {
	var postGreConn, err = postgresConn()
	PostgresConnection = postGreConn
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("PostgreSQL Connection has been initiated!")
}
