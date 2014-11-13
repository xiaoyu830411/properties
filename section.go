package properties

type Section interface {
	Id() string
	Get(key string) (string, error)
	Set(key string, value string) error
	Remove(key string) (string, error)
}


type mySection struct {
	id string
	pairs map[string]string
}

func newSection(id string) Section {
	return &mySection{id: id, pairs: make(map[string]string)}
}


func (this mySection) Id() string {
	return this.id
}

func (this mySection) Get(key string) (string, error) {
	v, ok := this.pairs[key]
	if !ok {
		return "", _NON_EXISTS_
	}

	return v, nil
}

func (this *mySection) Set(key string, value string) error {
	if len(key) == 0 {
		return _NULL_KEY_
	}

	this.pairs[key] = value

	return nil
}

func (this *mySection) Remove(key string) (string, error) {
	v, ok := this.pairs[key]

	if !ok {
		return "", _NON_EXISTS_
	}

	delete(this.pairs, key)

	return v, nil
}
