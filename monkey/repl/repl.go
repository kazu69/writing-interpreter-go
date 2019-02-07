package repl

import (
	"bufio"
	"fmt"
	"go/token"
	"io"

	"github.com/kazu69/golang-interpreter/monkey/lexer"
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

		// 改行が車で繰り返す
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Print("%+v\n", tok)
		}
	}
}
