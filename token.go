package properties

const (
	token_EOF = -1
	token_BLANK_LINE = 0
	token_SECTION_ID = 1
	token_KEY = 2
	token_VALUE = 3
)

var (
	token_names = [...]string {
		"EOF",
		"Blank Line",
		"Section ID",
		"Key",
		"Value",
	}
)

type token struct {
	Id int
	Name string
	Value string
}

func getTokenName(id int) string {
	return token_names[id + 1]
}


