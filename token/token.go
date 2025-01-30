// token/token.go

package token

type TokenType string

type Token struct {
	Type	TokenType 	// TokenType = string = many possible values for TokenType = distinguishing between types of tokens
	Literal string 		// may not perform same as int/byte, but simpler
}

const (
	ILLEGAL = "ILLEGAL" 	// token/char we don't know about
	EOF	= "EOF"		// "end of file" -> parser can stop

	// Identifiers + literals
	IDENT 	= "IDENT" 	// add, foobar, x, y...
	INT 	= "INT" 	// 1343456 

	// Operators
	ASSIGN	= "="
	PLUS 	= "+"
	MINUS	= "-"
	BANG	= "!"
	ASTERISK = "*"
	SLASH	= "/"
	
	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":	FUNCTION,
	"let": LET,
}

// checks keywords table for validation
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
