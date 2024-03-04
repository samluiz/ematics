package main

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENTIFIER = "IDENTIFIER"
	NUMBER = "NUMBER"

	ASSIGN = "="
	PLUS = "+"
	SLASH = "/"
	STAR = "*"
	MINUS = "-"

	LPAREN = "("
	RPAREN = ")"

	COMMA = ","
	SEMICOLON = ";"
	RETURN_ARROW = "->"

	FUNCTION = "FUNCTION"
)

var keywords = map[string]TokenType{
	"f": FUNCTION,
}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENTIFIER
}