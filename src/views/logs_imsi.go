package views

import (
	"fmt"
	"net/http"
	"strings"
)

func Logs_imsi(w http.ResponseWriter, r *http.Request) {
	var imsi = r.URL.Path
	imsi = strings.Split(imsi, "/api/logs/")[1] //ok, got imsi
	var count = r.URL.Query().Get("count")
	fmt.Printf("%s\n", count) //ok getting count
}
