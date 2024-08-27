package main

import (
	"log"
	"net/http"
	"rest-docker/database"
	"rest-docker/handler/resthttp"
	"rest-docker/pkg/record"
	"rest-docker/services"
)

func main() {
	// Init DB
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Init Resources
	recordPkg := record.New(db)

	// Init Services
	recordService, err := services.NewRecordService(services.RecordDependencies{
		DB:     db,
		Record: recordPkg,
	})
	if err != nil {
		log.Fatalf("Record service failed to initiate: %v", err)
	}

	r := resthttp.NewRouter(resthttp.RouterDependencies{
		RS: recordService,
	})

	port := "8080"

	log.Println("Service started at port " + port)
	http.ListenAndServe(":"+port, r)
}
