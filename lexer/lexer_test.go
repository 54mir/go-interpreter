package lexer

import (
	"go-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);

**! /<3>
return 9

!= ==


`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.ASTERISK, "*"},
		{token.ASTERISK, "*"},
		{token.BANG, "!"},
		{token.SLASH, "/"},
		{token.LT, "<"},
		{token.INT, "3"},
		{token.GT, ">"},
		{token.RETURN, "return"},
		{token.INT, "9"},
		{token.NOT_EQ, "!="},
		{token.EQ, "=="},
		{token.EOF, ""},
	}

	lex := New(input)
	for i, test := range tests {
		tok := lex.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("wrong type for test case: %d. expected: %+v, got: %+v for literal: %+v", i, test.expectedType, tok.Type, tok.Literal)
		}

		if token.TokenType(tok.Literal) != token.TokenType(test.expectedLiteral) {
			t.Fatalf("wrong literal for test case: %d. expected: '%+v', got: '%+v'", i, test.expectedLiteral, tok.Literal)
		}

	}

}
