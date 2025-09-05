package lexer

import (
	"testing"

	"github.com/Mr-xiaotian/Celestial-Language/Monkey/token"
)

type tokenCase struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func assertTokens(t *testing.T, input string, wants []tokenCase) {
	t.Helper() // 报错时定位到调用处
	l := New(input)

	for i, w := range wants {
		tok := l.NextToken()
		if tok.Type != w.expectedType || tok.Literal != w.expectedLiteral {
			t.Fatalf("case[%d] mismatch:\n  want: {type=%q, literal=%q}\n  got : {type=%q, literal=%q}",
				i, w.expectedType, w.expectedLiteral, tok.Type, tok.Literal)
		}
	}

	// 如果你希望在遇到错误就立即终止，可把上面的 t.Errorf 换成 t.Fatalf；
	// 但一般调 lexer 更希望一次看完所有不匹配，故保留 t.Errorf。
}

func TestNextToken0(t *testing.T) {
	input := `=+(){},;`

	tests := []tokenCase{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	assertTokens(t, input, tests)

}
func TestNextToken1(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	`

	tests := []tokenCase{
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
		{token.EOF, ""},
	}

	assertTokens(t, input, tests)

}

func TestNextToken2(t *testing.T) {
	input := `
	!-/*5;
	5 < 10 > 5;`

	tests := []tokenCase{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	assertTokens(t, input, tests)
}

func TestNextToken3(t *testing.T) {
	input := `
	if (5 < 10) {
		return true;
	} else {
		return false;
	}`

	tests := []tokenCase{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	assertTokens(t, input, tests)
}

func TestNextToken4(t *testing.T) {
	input := `
	10 == 10;
	10 != 9;
	`

	tests := []tokenCase{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	assertTokens(t, input, tests)
}

func TestNextToken5(t *testing.T) {
	input := `
	let s0 = "hello";
	let s1 = "world
	!";
	`

	tests := []tokenCase{
		{token.LET, "let"},
		{token.IDENT, "s0"},
		{token.ASSIGN, "="},
		{token.STRING, "hello"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "s1"},
		{token.ASSIGN, "="},
		{token.STRING, "world\n\t!"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	assertTokens(t, input, tests)
}
