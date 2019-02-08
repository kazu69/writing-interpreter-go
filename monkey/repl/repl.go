package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kazu69/golang-interpreter/monkey/lexer"
	"github.com/kazu69/golang-interpreter/monkey/token"
)

const PROMPT = ">"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// 改行がまでで繰り返す
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Print("%+v\n", tok)
		}
	}
}
