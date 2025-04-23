package migrations

import (
	// "fmt"
	// jazidb "httpproj1/db"
	"httpproj1/auth"
	"httpproj1/initializers"
	"httpproj1/logger"
	"httpproj1/shop"
	"log/slog"

	_ "github.com/lib/pq"
)

var myslog *slog.Logger = logger.GetLogger()

func SetUp() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		myslog.Error("Error in Load Config File!")
	}
	initializers.ConnectDB(&config)
}

func RunMigrations() {
	SetUp()

	initializers.DB.AutoMigrate(
		&shop.Brand{},
		&shop.Category{},
		&shop.Product{},
		&auth.User{},
	)
	myslog.Info("Migration Complete")
}

// func CreateTables() {
// 	myslog := logger.GetLogger()
// 	globalDatabase := jazidb.GetDatabase()

// 	myslog.Info("CREATING TABELS...")
// 	res, err := globalDatabase.Query(`create table if not exists Product (
// 	id INT PRIMARY KEY,
// 	title TEXT,
// 	price INT
// 	)`)
// 	if err != nil {
// 		myslog.Error(err.Error())
// 	} else {
// 		fmt.Println(res)
// 	}
// }
