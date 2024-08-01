package main

import (
	"fmt"
	"github.com/mrgyan89/redis-client-app/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to golang redis client application ")
	router := router.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Started server on port 4000")

}
