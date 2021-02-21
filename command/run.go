package command

import (
	"bytes"
	"errors"

	"github.com/lusingander/monkey/evaluator"
	"github.com/lusingander/monkey/lexer"
	"github.com/lusingander/monkey/object"
	"github.com/lusingander/monkey/parser"
)

func run(input string) error {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		return buildParserError(p.Errors())
	}

	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	evaluator.DefineMacros(program, macroEnv)
	expanded := evaluator.ExpandMacros(program, macroEnv)

	evaluated := evaluator.Eval(expanded, env)
	if errObj, ok := evaluated.(*object.Error); ok {
		return buildEvaluateError(errObj)
	}
	return nil
}

func buildParserError(errs []string) error {
	var out bytes.Buffer
	out.WriteString("ERROR:\n")
	for _, msg := range errs {
		out.WriteString("\t")
		out.WriteString(msg)
		out.WriteString("\n")
	}
	return errors.New(out.String())
}

func buildEvaluateError(err *object.Error) error {
	return errors.New(err.Inspect())
}
