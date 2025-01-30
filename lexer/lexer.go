package lexer

import "monkey/token"

type Lexer struct {
	input		string
	position	int // the position in the input (points to char)
	readPosition	int // current reading position in input (after current char)
	ch		byte // char being parsed
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar gives us next character and advances position in the input string
// TODO: To fully support Unicode & UTF-8, need to change l.ch from `byte` to `rune`,
// and change the way we read the next characters, since they could be multiple bytes wide
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" char
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition // always -> position last read
	l.readPosition += 1 // always -> next position
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	
	l.skipWhitespace()

	switch l.ch {
	case '=':
	tok = newToken(token.ASSIGN, l.ch)
	case '+':
	tok = newToken(token.PLUS, l.ch)
	case '-':
	tok = newToken(token.MINUS, l.ch)
	case '!':
	tok = newToken(token.BANG, l.ch)
	case '/':
	tok = newToken(token.SLASH, l.ch)
	case '*':
	tok = newToken(token.ASTERISK, l.ch)
	case '<':
	tok = newToken(token.LT, l.ch)
	case '>':
	tok = newToken(token.GT, l.ch)
	case ';':
	tok = newToken(token.SEMICOLON, l.ch)
	case ',':
	tok = newToken(token.COMMA, l.ch)
	case '(':
	tok = newToken(token.LPAREN, l.ch)
	case ')':
	tok = newToken(token.RPAREN, l.ch)
	case '{':
	tok = newToken(token.LBRACE, l.ch)
	case '}':
	tok = newToken(token.RBRACE, l.ch)
	case 0:
	tok.Literal = ""
	tok.Type = token.EOF
	
	// define default branch, checks for indentifiers if l.char unrecognized
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch) // create our ILLEGAL tokens
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// reads an identifier, then advances position until non-letter char
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// helper function, adds '_' so we can use var names like foo_bar... 
// NOTE: '!' and '?' are also possible
func isLetter (ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// eatWhitespace / consumeWhitespace
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// same as readIdentifier
// very simplified.
// TODO: floats? hex notation? octal notation?
// for now, Monkey doesn't support this :)
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// like isLetter, but returns whether passed byte is a Latin digit 0-9
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
