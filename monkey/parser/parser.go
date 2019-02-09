package parser

import (
	"github.com/kazu69/golang-interpreter/monkey/ast"
	"github.com/kazu69/golang-interpreter/monkey/lexer"
	"github.com/kazu69/golang-interpreter/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	cutToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// 2つのtokenを読み込む
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.cutToken = p.peekToken      // 現在のトークン
	p.peekToken = p.l.NextToken() // 次のトークン
}

func (p *Parser) ParseProgram() ast.Program {
	return nil
}
