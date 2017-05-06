package main

import (
	"log"
	"net/http"
	"utils"
	"views"
)

func main() {
	log.Println("Api server starting")
	utils.InitDb()
	views.Parser()
	http.HandleFunc("/api/logs/", views.Logs_imsi)
	http.HandleFunc("/", views.Dashboard)
	log.Println("Routing ready")
	log.Println("Setting up server")
	log.Println("Server up and running localhost:8080")
	http.ListenAndServe(":8080", nil)
}
