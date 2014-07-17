// the engine package will hold the data structures that meowy uses
// to represent people and things that people want to do
package engine

import (
	"errors"
	"log"
	"regexp"
)

// TODO better name than 'person'?
type Person struct {
	Things []string
}

// TODO find more replacements? ask amber
var replacements = map[string]string{
	"my":        "your",
	"I":         "you",
	"me":        "you",
	"mine":      "yours",
	"myself":    "yourself",
	"we":        "you",
	"us":        "you all",
	"ourselves": "yourselves",
	"our":       "your",
	"ours":      "yours",
}

// TODO name this better
func replaceInput(s string) string {
	for old, new := range replacements {
		re, err := regexp.Compile("(^|\\W)" + old + "(\\W|$)")
		if err != nil {
			log.Print(err)
			continue
		}
		s = re.ReplaceAllString(s, "${1}"+new+"${2}")
	}
	return s
}

func NewPerson() *Person {
	return &Person{
		Things: make([]string, 0, 1),
	}
}

// Pushes t to the front of Person.Things
func (p *Person) AddThing(t string) error {
	replaced := replaceInput(t)
	if replaced == "" {
		return errors.New("Bad input ):")
	}
	p.Things = append([]string{replaced}, p.Things...)
	return nil
}

// For template processing, we need to define the following methods:
// TopPriority and Rest. Then can then be used in the template as
// {{ .TopPriority }} and {{ range .Rest }}
func (p *Person) TopPriority() string {
	if len(p.Things) == 0 {
		return "________"
	}
	return p.Things[0]
}

func (p *Person) Rest() []string {
	if len(p.Things) < 2 {
		return nil
	}
	return p.Things[1:]
}

func (p *Person) Delete(i int) {
	if len(p.Things) == 0 || i < 0 || i >= len(p.Things) {
		// this shouldn't happen
		return
	}
	if i == len(p.Things)-1 {
		p.Things = p.Things[:i]
		return
	}
	p.Things = append(p.Things[:i], p.Things[i+1:]...)
}

func (p *Person) Promote(i int) {
	if len(p.Things) == 0 || i < 0 || i >= len(p.Things) {
		// this shouldn't happen
		return
	}
	front := p.Things[0]
	p.Things[0] = p.Things[i]
	p.Things[i] = front
}
