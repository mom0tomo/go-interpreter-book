package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mom0tomo/go-interpreter-book/monkey/lexer"
	"github.com/mom0tomo/go-interpreter-book/monkey/parser"
)
const PROMPT = ">> "

const GOPHER_FACE = `
ʕ ◔ϖ◔ʔ
ʕ ◔ϖ◔ʔʕ ◔ϖ◔ʔ
ʕ ◔ϖ◔ʔʕ ◔ϖ◔ʔʕ ◔ϖ◔ʔ
`

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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, GOPHER_FACE)
	io.WriteString(out, "Woops! Web ran into some gopher business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
