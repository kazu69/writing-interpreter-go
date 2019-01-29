package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGSL = "ILLEGSL"
	EOF     = "EOF"

	// 識別子・リテラル
	IDEN = "IDEN"
	INT  = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
