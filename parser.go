package properties

import (
	"fmt"
)

type parser struct {
	lexer lexer
	t token

	elements map[string]string
	sections map[string](map[string]string)
}

func newParser(lexer lexer) parser {
	parser := parser{
		lexer: lexer,
		elements: make(map[string]string),
		sections: make(map[string](map[string]string)),
	}

	parser.consume()

	return parser
}

func (this *parser) properties()  {
	for this.t.Id != token_EOF {
		switch this.t.Id {
			case token_BLANK_LINE: this.consume()
			case token_SECTION_ID: this.section()
			case token_KEY : this.element(&this.elements)
			default:
				panic(fmt.Sprintf("Invalid Token[%+v]", this.t))
		}
	}
}

func (this *parser) section() {
	id :=this.match(token_SECTION_ID)

	if _, ok := this.sections[id.Value]; ok {
		panic(fmt.Sprintf("The section[%v] has been duplication", id.Value))
	}

	section := make(map[string]string)
	this.sections[id.Value] = section

	for this.t.Id != token_BLANK_LINE && this.t.Id != token_EOF && this.t.Id != token_SECTION_ID {
		this.element(&section)
	}

}

func (this *parser) element(section *map[string]string) {
	key := this.match(token_KEY)
	value := this.match(token_VALUE)

	if _, ok := (*section)[key.Value]; ok {
		panic(fmt.Sprintf("The key[%v] has been duplication", key.Value))
	}

	(*section)[key.Value] = value.Value
}

func (this *parser) match(token int) (t token) {
	if this.t.Id == token {
		t = this.t
		this.consume()
		return
	}

	panic(fmt.Sprintf("Expect a %v", getTokenName(token)))
}

func (this *parser) consume() {
	this.t = this.lexer.NextToken()
}
