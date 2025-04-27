package initializers

import (
	"context"
	"fmt"
	"httpproj1/logger"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *gorm.DB
var MONGO *mongo.Client

var myslog *slog.Logger = logger.GetLogger()

var ProductCollection *mongo.Collection

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		myslog.Error("Failed to connect to the Postgres Database")
	}
	fmt.Println("🚀 Connected Successfully to the Postgres Database")

	MONGO, err = mongo.Connect(options.Client().ApplyURI(config.MongoDbUri))
	if err != nil {
		myslog.Error("❌ Failed to connect to the Mongo Database:", err)
		panic(err)
	}
	pingErr := MONGO.Ping(context.TODO(), nil)
	if pingErr != nil {
		myslog.Error("❌ Ping failed. Could not connect to MongoDB:", pingErr)
	}
	// defer func() {
	// 	if err := MONGO.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	fmt.Println("🚀 Connected Successfully to the Mongo Database")

	ProductCollection = MONGO.Database("db").Collection("product")
}
