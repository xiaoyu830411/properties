package properties

type Section interface {
	Id() string
	Get(key string) (string, bool)
	Set(key string, value string) error
	Remove(key string) (string, error)
}


type mySection struct {
	id string
	elements map[string]string
}

func newSection(id string) Section {
	return &mySection{id: id, elements: make(map[string]string)}
}


func (this mySection) Id() string {
	return this.id
}

func (this mySection) Get(key string) (string, bool) {
	v, ok := this.elements[key]
	return v, ok
}

func (this *mySection) Set(key string, value string) error {
	if len(key) == 0 {
		return _NULL_KEY_
	}

	this.elements[key] = value

	return nil
}

func (this *mySection) Remove(key string) (string, error) {
	v, ok := this.elements[key]

	if !ok {
		return "", _NON_EXISTS_
	}

	delete(this.elements, key)

	return v, nil
}
