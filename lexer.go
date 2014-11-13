package properties

import(
	"strings"
	"fmt"
)

const (
	__EOF__ = rune(-1)
)

type lexer struct {
	stream []rune
	p int
	c rune
}

func newLexer(content string) lexer {
	l := lexer{stream: []rune(content), p: -1}
	l.consume()

	return l
}

func (this *lexer) NextToken() token {
	for this.c != __EOF__ {
		switch this.c {
			case ' ', '\t', '\r': this.ws()
			case '#': this.comments()
			case '[': return this.sectionIdToken()
			case '=': return this.valueToken()
			case '\n': return this.blankLineToken()
			default: return this.keyToken()
		}
	}

	return token{Id: token_EOF, Name: getTokenName(token_EOF), Value: "-1"}
}

func (this *lexer) comments() {
	this.match('#')

	for this.isNotEnd() {
		this.consume()
	}

	this.consume()
}

func (this *lexer) sectionIdToken() token {
	this.match('[')

	id := make([]rune, 0, 256)
	for this.c != ']' && this.isNotEnd() {
		id = append(id, this.c)
		this.consume()
	}

	this.match(']')

	this.ws()

	if this.c == '\r' {
		this.consume()
	}

	if this.isNotEnd() {
		panic("The section must be end by new line, or be end of file!")
	} else {
		this.consume()
	}

	s := string(id)
	if len(s) == 0 {
		panic("The section Id cannot be null!")
	}

	return token{Id: token_SECTION_ID, Name: getTokenName(token_SECTION_ID), Value: s}
}

func (this *lexer) keyToken() token {
	key := make([]rune, 0, 256)
	for this.c != '=' && this.c != '\r' && this.isNotEnd() {
		key = append(key, this.c)
		this.consume()
	}

	if !this.isNotEnd() {
		this.consume()
	}

	return token{Id: token_KEY, Name: getTokenName(token_KEY), Value: strings.TrimSpace(string(key))}
}

func (this *lexer) valueToken() token {
	this.match('=')

	value := make([]rune, 0, 256)

	for this.c != '\r' && this.isNotEnd() {
		value = append(value, this.c)
		this.consume()
	}

	if this.c == '\r' {
		this.consume()
	}

	if this.isNotEnd() {
		panic("The key = value must be end by new line, or be end of file!")
	} else {
		this.consume()
	}

	s := string(value)
	s = strings.TrimSpace(s)

	return token{Id: token_VALUE, Name: getTokenName(token_VALUE), Value: s}
}

func (this *lexer) blankLineToken() token {
	this.match('\n')

	return token{Id: token_BLANK_LINE, Name: getTokenName(token_BLANK_LINE), Value: "\n"}
}


func (this *lexer) match(c rune) {
	if this.c == c {
		this.consume()
	} else {
		panic(fmt.Sprintf("Expect a '%v' on %v", string(c), this.p))
	}
}

func (this *lexer) consume() {
	this.p++

	if this.p >= len(this.stream) {
		this.c = __EOF__
	} else {
		this.c = this.stream[this.p]
	}

}

func (this *lexer) ws() {
	for this.c == ' ' || this.c == '\t' {
		this.consume()
	}
}

func (this *lexer) isNotEnd() bool {
	return  (this.c != '\n' && this.c != __EOF__)
}
