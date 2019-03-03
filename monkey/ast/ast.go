package ast

import "github.com/kazu69/golang-interpreter/monkey/token"

// ast.Program [Statements]
//  |
// ast.LetStatement
//  |- Name ---- ast.Identifier
//  |- Valie --- ast.Expression

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token // token.Ident トークン
	Value string
}

type Expression interface {
	Node
	expressionNode()
}

type LetStatement struct {
	Token token.Token // token>LET トークン
	Name  *Identifier // 識別子 ast.Identifier
	Value Expression  // 値    ast.Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
