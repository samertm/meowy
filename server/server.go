package server

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/samertm/meowy/engine"
	"github.com/samertm/samerhttp/form"
	"github.com/samertm/samerhttp/session"
)

// Currently assigns a new Person to every session. I'm not sure if
// I should only create new Person types when the user submits a
// change request.
func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Print(err)
			return
		}
		v, ok := s.CookieGet(r)
		if !ok {
			v = &engine.Person{}
		}
		err = t.Execute(w, v.(*engine.Person))
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
		v, ok := s.CookieGet(r)
		if !ok {
			v = engine.NewPerson()
			s.CookieSet(r, v)
		}
		p, ok := v.(*engine.Person)
		if !ok {
			log.Println("Didn't store a person D:")
		}
		p.AddThing(f["thing"][0])
	}
}

func handleThingDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		f, err := form.Parse(r, "delete")
		if err != nil {
			log.Print(err)
			return
		}
		err = withPerson(r, func(e *engine.Person) {
			i, err := strconv.Atoi(f["delete"][0])
			if err != nil {
				log.Print("expected an int")
				return
			}
			e.Delete(i)
		})
		if err != nil {
			log.Print(err)
		}
	}
}

func handleThingPromote(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		f, err := form.Parse(r, "promote")
		if err != nil {
			log.Print(err)
			return
		}
		err = withPerson(r, func(e *engine.Person) {
			i, err := strconv.Atoi(f["promote"][0])
			if err != nil {
				log.Print("expected an int")
				return
			}
			e.Promote(i)
		})
		if err != nil {
			log.Print(err)
		}
	}
}

func withPerson(r *http.Request, fn func(*engine.Person)) error {
	v, ok := s.CookieGet(r)
	if !ok {
		return errors.New("auth cookie not set")
	}
	p, ok := v.(*engine.Person)
	if !ok {
		return errors.New("cookie set to invalid type")
	}
	fn(p)
	return nil
}

var s = session.New()
var defaultThingStruct = struct{ Thing string }{Thing: "________"}

func ListenAndServe(ip, prefix string) {
	r := mux.NewRouter()
	r.HandleFunc(prefix+"/", handleIndex)
	r.HandleFunc(prefix+"/thing/change", handleThingChange)
	r.HandleFunc(prefix+"/thing/delete", handleThingDelete)
	r.HandleFunc(prefix+"/thing/promote", handleThingPromote)
	r.PathPrefix(prefix + "/").Handler(
		http.StripPrefix(prefix+"/", http.FileServer(http.Dir("./static/"))))
	// I don't think I need to append prefix to the front of "/"
	http.Handle("/", r)
	err := http.ListenAndServe(ip, nil)
	if err != nil {
		log.Print(err)
	}
}
