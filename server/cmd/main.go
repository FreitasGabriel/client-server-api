package main

import (
	"log"
	"net/http"

	"github.com/FreitasGabriel/client-server-api/server/api/controllers"
	"github.com/FreitasGabriel/client-server-api/server/config"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Starting server")
	mux := http.NewServeMux()

	database, err := config.InitDatabase()
	if err != nil {
		log.Println("error to init database", err)
		return
	}

	ctr := controllers.NewQuotationInterface(database)

	mux.HandleFunc("/cotacao", ctr.MakeQuotation)
	http.ListenAndServe(":8080", mux)
}
