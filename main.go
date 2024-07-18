package main

import (
	"go-template/config"
	"go-template/routers"
	"log"
	"net/http"
)

func main() {
	config.Connect()

	router := routers.InitRouter()

	log.Println("server di port 8080 bro")
	log.Fatal(http.ListenAndServe(":8080", router))
}
