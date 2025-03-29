package migrations

import (
	"fmt"
	jazidb "httpproj1/db"
	"httpproj1/logger"

	_ "github.com/lib/pq"
)

func CreateTables() {
	myslog := logger.GetLogger()
	globalDatabase := jazidb.GetDatabase()

	myslog.Info("CREATING TABELS...")
	res, err := globalDatabase.Query(`create table if not exists Product (
	id INT PRIMARY KEY,
	title TEXT,
	price INT
	)`)
	if err != nil {
		myslog.Error(err.Error())
	} else {
		fmt.Println(res)
	}
}
