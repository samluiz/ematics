package main

type Lexer struct {
	input        string
	position     int
	nextPosition int
	char 			 byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.nextCharacter()
	return l
}

// this function basically moves to the next character and stops if the next character is the end of the input
func (l *Lexer) nextCharacter() {
	if l.nextPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

// this function returns what the next char will be
// but don't increment the read position
func (l *Lexer) peekNextChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

// this function, well, skips whitespaces
func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.nextCharacter()
	}
}

// this function returns a token with the type and the literal
func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// what the lexer will identify as letter. this allows variable names
// like 'foo_bar'
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// what the lexer will identify as a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// this function reads the identifier name and returns the token
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.nextCharacter()
	}
	return l.input[position:l.position]
}

// this function reads the identifier name and returns the token
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.nextCharacter()
	}
	return l.input[position:l.position]
}

// this function reads the next token
// this function will be called in a loop until the end of the input
func (l *Lexer) ReadNextToken() Token {
	var token Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		token = newToken(ASSIGN, l.char)
	case ';':
		token = newToken(SEMICOLON, l.char)
	case '(':
		token = newToken(LPAREN, l.char)
	case ')':
		token = newToken(RPAREN, l.char)
	case ',':
		token = newToken(COMMA, l.char)
	case '+':
		token = newToken(PLUS, l.char)
	case '-':
		if l.peekNextChar() == '>' {
			l.nextCharacter()
			token = Token{Type: RETURN_ARROW, Literal: "->"}
		} else {
		token = newToken(MINUS, l.char)
		}
	case '*':
		token = newToken(STAR, l.char)
	case '/':
		token = newToken(SLASH, l.char)
	case 0:
		token.Literal = ""
		token.Type = EOF
	default:
		if isLetter(l.char) {
			token.Literal = l.readIdentifier()
			token.Type = LookupIdent(token.Literal)
			return token
		} else if isDigit(l.char) {
			token.Type = NUMBER
			token.Literal = l.readNumber()
			return token
		} else {
			token = newToken(ILLEGAL, l.char)
		}
	}

	l.nextCharacter()
	return token
}