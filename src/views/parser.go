package views

import (
	"html/template"
	"io/ioutil"
	"log"
)

var Templates map[string]*template.Template

func Parser() {
	log.Println("Caching templates")
	Templates = make(map[string]*template.Template)
	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename := file.Name()
		t := template.New(filename)
		t, err = t.ParseFiles("./templates/"+filename, "./templates/base.html")
		Templates[filename] = t
	}
	for key := range Templates {
		log.Println("Found", key, "template")
	}
	log.Println("Total templates", len(Templates))
}
