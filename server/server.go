package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samertm/samerhttp/form"
	"github.com/samertm/samerhttp/session"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Print(err)
			return
		}
		v, ok := s.Get("onesession")
		if !ok {
			v = struct{ Thing string }{Thing: "________"}
		}
		err = t.Execute(w, v)
		if err != nil {
			log.Print(err)
		}
	}
}

func handleThingChange(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		f, err := form.Parse(r, "thing")
		if err != nil {
			log.Print(err)
			return
		}
		s.Set("onesession", struct{ Thing string }{Thing: f["thing"][0]})
	}
}

// TODO deal with more than one sesion
var s = session.New()

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
