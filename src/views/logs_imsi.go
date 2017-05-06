package views

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"utils"
)

func Logs_imsi(w http.ResponseWriter, r *http.Request) {
	var imsi = r.URL.Path
	imsi = strings.Split(imsi, "/api/logs/")[1] //ok, got imsi
	var count = r.URL.Query().Get("count")      //ok getting count
	var db = utils.GetDb()
	results, err := db.Query("SELECT field,value from eos_log where imsi=? LIMIT ?", "204043255462105", 10)
	if err != nil {
		log.Fatal("Db query")
	}
	defer results.Close()
	var value string
	var field string
	for results.Next() {
		err := results.Scan(&field, &value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s : %s\n", field, value)
	}
	fmt.Printf("%d\n", count)
}
