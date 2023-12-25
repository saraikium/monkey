package repl

import (
	"bufio"
	"fmt"
	"github.com/saraikium/monkey/lexer"
	"github.com/saraikium/saraikium/monkey/token"
	"io"
)

const PROMPT = ">>"

func start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {

		}
	}

}