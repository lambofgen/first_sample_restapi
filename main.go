package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"geon/api3/dbprovider"
	"geon/api3/router"
)

const port = "5055"

func main() {

	db := dbprovider.InitDatabase()
	defer func() {
		log.Print("destroy database")
		db.Close()
		log.Println("============ EXIT")
	}()

	m := mux.NewRouter().StrictSlash(true)
	router.Mount(m, db)

	log.Println("Start server api on port: ", port)
	log.Println("============ START")
	log.Fatal(http.ListenAndServe(":"+port, m))
}
