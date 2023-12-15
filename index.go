package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	User  string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("index.html"))
		Pages := map[string][]Page{
			"pages": {
				{Title: "dashboard", User: "ahmed"},
				{Title: "profile", User: "aussama"},
				{Title: "settings", User: "ghanama"},
				{Title: "logout", User: "admin"},
				{Title: "login", User: "ahmed"},
				{Title: "register", User: "ghanama"},
			},
		}
		tmp.Execute(w, Pages)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		user := r.PostFormValue("user")
		html := fmt.Sprintf(`<li class="shadow-md text-center p-2 w-full rounded-lg">%s - %s </li>`, title, user)
		temp , _ :=  template.New("t").Parse(html)
		temp.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add", h2)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
