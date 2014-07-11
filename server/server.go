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
		c, err := r.Cookie("sessionid")
		// TODO look into the structure of this?
		var v interface{}
		if err != nil {
			v = defaultThingStruct
		} else {
			var ok bool // declare here so we don't shadow 'v'
			v, ok = s.Get(c.Value)
			if !ok {
				v = defaultThingStruct
			}
		}
		err = t.Execute(w, v.(struct{ Thing string }))
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
		c, err := r.Cookie("sessionid")
		if err != nil {
			log.Println("No cookie set")
			return
		}
		s.Set(c.Value, struct{ Thing string }{Thing: f["thing"][0]})
	}
}

// TODO deal with more than one sesion
var s = session.New()
var defaultThingStruct = struct{ Thing string }{Thing: "________"}

func ListenAndServe(ip, prefix string) {
	r := mux.NewRouter()
	r.HandleFunc(prefix+"/", handleIndex)
	r.HandleFunc(prefix+"/thing/change", handleThingChange)
	r.PathPrefix(prefix + "/").Handler(
		http.StripPrefix(prefix+"/", http.FileServer(http.Dir("./static/"))))
	// I don't think I need to append prefix to the front of "/"
	http.Handle("/", r)
	err := http.ListenAndServe(ip, nil)
	if err != nil {
		log.Print(err)
	}
}
