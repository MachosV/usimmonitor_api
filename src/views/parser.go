package views

import (
	"fmt"
	"html/template"
	"io/ioutil"
)

var Templates map[string]*template.Template

func Parser() {
	fmt.Println("[*]Caching templates")
	Templates = make(map[string]*template.Template)
	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		filename := file.Name()
		t := template.New(filename)
		t, err = t.ParseFiles("./templates/"+filename, "./templates/base.html")
		Templates[filename] = t
	}
	for key := range Templates {
		fmt.Println("[*]Found", key, "template")
	}
	fmt.Println("[*]Total templates", len(Templates))
}
