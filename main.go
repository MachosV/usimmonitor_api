package main

import (
	"fmt"
	"net/http"
	"views"
)

func main() {
	views.Parser()
	http.HandleFunc("/api/logs/", views.Logs_imsi)
	http.HandleFunc("/", views.Dashboard)
	fmt.Println("[*]Routing ready")
	fmt.Println("[*]Setting up server")
	fmt.Println("[*]Server up and running, localhost:8080")
	http.ListenAndServe(":8080", nil)
}
