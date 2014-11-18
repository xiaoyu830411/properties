package properties

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Properties interface {
	Get(key string) (string, bool)
	Set(key string, value string) error
	Remove(key string) (string, error)

	Elements() map[string]string

	GetSection(id string) (Section, bool)
	SetSection(section Section) error
	RemoveSection(id string) (Section, error)

	Sections() map[string]Section
}

type myProperties struct {
	elements map[string]string
	sections map[string]Section
}

func Load(file string) (Properties, error) {
	elements, sections, err := parseFile(file)
	if err != nil {
		return nil, err
	}

	return newProperties(elements, sections), nil
}

func newProperties(elements map[string]string, sects map[string](map[string]string)) Properties {
	var sections = make(map[string]Section)
	for id, values := range sects {
		section := newSection(id)
		sections[id] = section

		for key, value := range values {
			section.Set(key, value)
		}
	}

	return &myProperties{elements: elements, sections: sections}
}

func NewProperties() Properties {
	return &myProperties{elements: make(map[string]string), sections: make(map[string]Section)}
}

func parseFile(file string) (elements map[string]string, sections map[string](map[string]string), err error) {
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, nil, err
	}

	return parse(bytes)
}

func parse(bytes []byte) (elements map[string]string, sections map[string](map[string]string), err error) {
	defer func() {
		r := recover()

		if r != nil {
			switch x := r.(type) {
			case error:
				err = x
			default:
				err = errors.New(fmt.Sprintf("%v", x))
			}

			elements = nil
			sections = nil
		}
	}()

	lexer := newLexer(string(bytes))
	parser := newParser(lexer)
	parser.properties()

	elements = parser.elements
	sections = parser.sections

	return
}

func (this myProperties) Get(key string) (string, bool) {
	v, ok := this.elements[key]
	return v, ok
}

func (this *myProperties) Set(key string, value string) error {
	if len(key) == 0 {
		return _NULL_KEY_
	}

	this.elements[key] = value

	return nil
}

func (this *myProperties) Remove(key string) (string, error) {
	v, ok := this.elements[key]

	if !ok {
		return "", _NON_EXISTS_
	}

	delete(this.elements, key)

	return v, nil
}

func (this myProperties) Elements() map[string]string {
	return this.elements
}

func (this myProperties) GetSection(id string) (Section, bool) {
	v, ok := this.sections[id]
	return v, ok
}

func (this *myProperties) SetSection(section Section) error {
	if section.Id() == "" {
		return _NULL_KEY_
	}

	this.sections[section.Id()] = section

	return nil
}

func (this *myProperties) RemoveSection(id string) (Section, error) {
	v, ok := this.sections[id]

	if !ok {
		return nil, _NON_EXISTS_
	}

	delete(this.sections, id)

	return v, nil
}

func (this myProperties) Sections() map[string]Section {
	return this.sections
}
