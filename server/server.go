package server

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samertm/samerhttp/form"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Print(err)
			return
		}
		err = t.Execute(w, struct{ Thing string }{Thing: "yo"})
		if err != nil {
			log.Print(err)
		}
	}
}

func handleThingChange(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, err := form.Parse(r, "thing")
		if err != nil {
			log.Print(err)
			return
		}
	}
}

func ListenAndServe(ip string) {
	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/thing/change", handleThingChange)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)
	err := http.ListenAndServe(ip, nil)
	if err != nil {
		log.Print(err)
	}
}
