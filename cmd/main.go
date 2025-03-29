package main

import (
	"httpproj1/apis"
	jazidb "httpproj1/db"
	"httpproj1/logger"
	"httpproj1/migrations"
	"net/http"
)

// var product *Product = &Product{Id: 1, Title: "jazi", Price: 100}

func main() {
	myslog := logger.GetLogger()
	//Routes
	http.HandleFunc("/", apis.GetRoot)
	http.HandleFunc("/create", apis.CreateProduct)

	//DATABASE
	db_connection := jazidb.ConnectDatabase()
	defer db_connection.Close()

	migrations.CreateTables()

	//SERVER LISTNER
	myslog.Info("Start Listening to port 5050")

	serverErr := http.ListenAndServe(":5050", nil)

	if serverErr != nil {
		myslog.Error("Somthing is Wrong")
	}
}
