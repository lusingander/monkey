package ast

import (
	"bytes"

	"github.com/lusingander/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (s *LetStatement) statementNode() {}

func (s *LetStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.Name.String())
	out.WriteString(" = ")
	if s.Value != nil {
		out.WriteString(s.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (s *ReturnStatement) statementNode() {}

func (s *ReturnStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(s.TokenLiteral() + " ")
	if s.ReturnValue != nil {
		out.WriteString(s.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // First token of the expression
	Expression Expression
}

func (s *ExpressionStatement) statementNode() {}

func (s *ExpressionStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *ExpressionStatement) String() string {
	if s.Expression != nil {
		return s.Expression.String()
	}
	return ""
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (l *IntegerLiteral) expressionNode() {}

func (l *IntegerLiteral) TokenLiteral() string {
	return l.Token.Literal
}

func (l *IntegerLiteral) String() string {
	return l.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token // prefix token, e.g. "!"
	Operator string
	Right    Expression
}

func (e *PrefixExpression) expressionNode() {}

func (e *PrefixExpression) TokenLiteral() string {
	return e.Token.Literal
}

func (e *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(e.Operator)
	out.WriteString(e.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token // infix token, e.g. "+"
	Left     Expression
	Operator string
	Right    Expression
}

func (e *InfixExpression) expressionNode() {}

func (e *InfixExpression) TokenLiteral() string {
	return e.Token.Literal
}

func (e *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString(" " + e.Operator + " ")
	out.WriteString(e.Right.String())
	out.WriteString(")")
	return out.String()
}