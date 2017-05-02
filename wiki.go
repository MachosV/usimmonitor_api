package main

import (
	"fmt"
	"net/http"
	"views"
)

func main() {
	views.Parser()
	http.HandleFunc("/", views.Dashboard)
	fmt.Println("[*]Routing ready")
	fmt.Println("[*]Setting up server")
	fmt.Println("[*]Server up and running")
	http.ListenAndServe(":8080", nil)

}
