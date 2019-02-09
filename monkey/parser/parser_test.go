package parser

import (
	"testing"

	"github.com/kazu69/golang-interpreter/monkey/ast"
	"github.com/kazu69/golang-interpreter/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatal("program.Statements does not contain 3 statements. get=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Error("s.TokenLiteral not `let` . got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Error("s not *ast.LetStatement. got=%q", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Error("letStmt.Name.Value not '%s'. got=%q", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Value.TokenLiteral() != name {
		t.Error("letStmt.Value.TokenLiteral() not '%s'. got=%q", name, letStmt.Value.TokenLiteral())
		return false
	}

	return true
}
