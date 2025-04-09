package jazidb

import (
	"database/sql"
	"fmt"
	"httpproj1/logger"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5444
	user     = "citizix_user"
	password = "S3cret"
	dbname   = "citizix_db"
)

var globalDatabase *sql.DB

func GetDatabase() *sql.DB {
	return globalDatabase
}

func ConnectDatabase() *sql.DB {
	myslog := logger.GetLogger()

	//DB
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		myslog.Info(fmt.Sprintf("Connected to database on port %d", port))
		globalDatabase = db
	}
	// defer db.Close()

	return db

}
