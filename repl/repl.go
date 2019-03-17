package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/NAKKA-K/learn-interpreter-in-go/lexer"
	"github.com/NAKKA-K/learn-interpreter-in-go/token"
)

// PROMPT is the waiting read icon of CUI interface
const PROMPT = ">> "

// Start prompt
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
