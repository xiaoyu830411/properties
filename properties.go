package properties

import (
	"io/ioutil"
	"errors"
	"fmt"
)

type Properties interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Remove(key string) (string, error)

	GetSection(id string) (Section, error)
	SetSection(section Section) error
	RemoveSection(id string) (Section, error)
}

type myProperties struct {
	pairs map[string]string
	sections map[string]Section
}

func Load(file string) (Properties, error) {
	pairs, sections, err := parseFile(file)
	if err != nil {
		return nil, err
	}

	return newProperties(pairs, sections), nil
}

func newProperties(pairs map[string]string, sects map[string](map[string]string)) Properties {
	var sections = make(map[string]Section)
	for id, values := range sects {
		section := newSection(id)
		sections[id] = section

		for key, value := range values {
			section.Set(key, value)
		}
	}

	return &myProperties{pairs: pairs, sections: sections}
}

func NewProperties() Properties {
	return &myProperties{pairs: make(map[string]string), sections: make(map[string]Section)}
}

func parseFile(file string) (pairs map[string]string, sections map[string](map[string]string), err error) {
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, nil, err
	}

	return parse(bytes)
}

func parse(bytes []byte) (pairs map[string]string, sections map[string](map[string]string), err error) {
	defer func() {
		r := recover()

		if r != nil {
			switch x := r.(type) {
				case error : err = x
				default : err = errors.New(fmt.Sprintf("%v", x))
			}

			pairs = nil
			sections = nil
		}
	} ()

	lexer := newLexer(string(bytes))
	parser := newParser(lexer)
	parser.properties()

	pairs = parser.pairs
	sections = parser.sections

	return
}

func (this myProperties) Get(key string) (string, error) {
	v, ok := this.pairs[key]
	if !ok {
		return "", _NON_EXISTS_
	}

	return v, nil
}

func (this *myProperties) Set(key string, value string) error {
	if len(key) == 0 {
		return _NULL_KEY_
	}

	this.pairs[key] = value

	return nil
}

func (this *myProperties) Remove(key string) (string, error) {
	v, ok := this.pairs[key]

	if !ok {
		return "", _NON_EXISTS_
	}

	delete(this.pairs, key)

	return v, nil
}

func (this myProperties) GetSection(id string) (Section, error) {
	v, ok := this.sections[id]
	if !ok {
		return nil, _NON_EXISTS_
	}

	return v, nil
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
