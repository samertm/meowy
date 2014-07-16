// the engine package will hold the data structures that meowy uses
// to represent people and things that people want to do
package engine

import "strings"

// TODO better name than 'person'?
type Person struct {
	Things []string
}

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

func NewPerson() *Person {
	return &Person{
		Things: make([]string, 0, 1),
	}
}

// Pushes t to the front of Person.Things
func (p *Person) AddThing(t string) {
	p.Things = append([]string{t}, p.Things...)
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

// Modifying PlaceholderPerson is illegal
var PlaceholderPerson = &Person{}
