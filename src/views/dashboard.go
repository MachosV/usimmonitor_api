package views

import (
	"net/http"
)

type Context struct {
	Data string
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	context := Context{Data: r.URL.Path[1:]}
	t, _ := Templates["dashboard.html"]
	t.ExecuteTemplate(w, "base", context)
}
