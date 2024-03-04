package main

import (
	"testing"
)

func TestLexer(t *testing.T) {
	input := `
	a = 2;
	b = 3;
	c = 4;

	add f(x, y, z) -> x + y * z / 2 - 1;

	add(a, b, c);
`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
}{
		{IDENTIFIER, "a"},
		{ASSIGN, "="},
		{NUMBER, "2"},
		{SEMICOLON, ";"},
		{IDENTIFIER, "b"},
		{ASSIGN, "="},
		{NUMBER, "3"},
		{SEMICOLON, ";"},
		{IDENTIFIER, "c"},
		{ASSIGN, "="},
		{NUMBER, "4"},
		{SEMICOLON, ";"},
		{IDENTIFIER, "add"},
		{FUNCTION, "f"},
		{LPAREN, "("},
		{IDENTIFIER, "x"},
		{COMMA, ","},
		{IDENTIFIER, "y"},
		{COMMA, ","},
		{IDENTIFIER, "z"},
		{RPAREN, ")"},
		{RETURN_ARROW, "->"},
		{IDENTIFIER, "x"},
		{PLUS, "+"},
		{IDENTIFIER, "y"},
		{STAR, "*"},
		{IDENTIFIER, "z"},
		{SLASH, "/"},
		{NUMBER, "2"},
		{MINUS, "-"},
		{NUMBER, "1"},
		{SEMICOLON, ";"},
		{IDENTIFIER, "add"},
		{LPAREN, "("},
		{IDENTIFIER, "a"},
		{COMMA, ","},
		{IDENTIFIER, "b"},
		{COMMA, ","},
		{IDENTIFIER, "c"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		token := lexer.ReadNextToken()

		if (token.Type != tt.expectedType) {
			t.Fatalf("tests[%d] - wrong token type. expected=%q, actual=%q", i, tt.expectedType, token.Type)
		}

		if (token.Literal != tt.expectedLiteral) {
			t.Fatalf("tests[%d] - wrong token literal. expected=%q, actual=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}