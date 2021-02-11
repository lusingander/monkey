package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/lusingander/monkey/evaluator"
	"github.com/lusingander/monkey/lexer"
	"github.com/lusingander/monkey/object"
	"github.com/lusingander/monkey/parser"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvronment()

	for {
		fmt.Fprintf(out, prompt)

		if !scanner.Scan() {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
