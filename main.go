package main

import (
	"fmt"
	"log"
	"net/http"

	"CronJob/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4002", r))
}
