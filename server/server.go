package server

import (
	"html/template"
	"log"
	"net/http"
	"strings"

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
		v, ok := s.CookieGet(r)
		if !ok {
			v = defaultThingStruct
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
		s.CookieSet(r, struct{ Thing string }{Thing: replaceInput(f["thing"][0])})
	}
}

// TODO deal with more than one sesion
var s = session.New()
var defaultThingStruct = struct{ Thing string }{Thing: "________"}

// TODO refactor into 'engine' package
// TODO find more replacements? ask amber
var replacements = map[string]string{
	"my": "your",
}

func replaceInput(s string) string {
	for old, new := range replacements {
		s = strings.Replace(s, old, new, -1)
	}
	return s
}

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
