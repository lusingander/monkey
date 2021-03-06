package ast

import (
	"bytes"
	"strings"

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

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (l *FloatLiteral) expressionNode() {}

func (l *FloatLiteral) TokenLiteral() string {
	return l.Token.Literal
}

func (l *FloatLiteral) String() string {
	return l.Token.Literal
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (l *StringLiteral) expressionNode() {}

func (l *StringLiteral) TokenLiteral() string {
	return l.Token.Literal
}

func (l *StringLiteral) String() string {
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

type IfExpression struct {
	Token       token.Token // token.IF
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (e *IfExpression) expressionNode() {}

func (e *IfExpression) TokenLiteral() string {
	return e.Token.Literal
}

func (e *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(e.Condition.String())
	out.WriteString(" ")
	out.WriteString(e.Consequence.String())
	if e.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(e.Alternative.String())
	}
	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (s *BlockStatement) statementNode() {}

func (s *BlockStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *BlockStatement) String() string {
	var out bytes.Buffer
	for _, stmt := range s.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // token.FUNCTION
	Parameters []*Identifier
	Body       *BlockStatement
}

func (l *FunctionLiteral) expressionNode() {}

func (l *FunctionLiteral) TokenLiteral() string {
	return l.Token.Literal
}

func (l *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := make([]string, 0)
	for _, p := range l.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(l.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(l.Body.String())
	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression // Identifier or FunctionLiteral
	Arguments []Expression
}

func (e *CallExpression) expressionNode() {}

func (e *CallExpression) TokenLiteral() string {
	return e.Token.Literal
}

func (e *CallExpression) String() string {
	var out bytes.Buffer
	args := make([]string, 0)
	for _, a := range e.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(e.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (l *ArrayLiteral) expressionNode() {}

func (l *ArrayLiteral) TokenLiteral() string {
	return l.Token.Literal
}

func (l *ArrayLiteral) String() string {
	var out bytes.Buffer
	elems := make([]string, 0)
	for _, e := range l.Elements {
		elems = append(elems, e.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")
	return out.String()
}

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (l *HashLiteral) expressionNode() {}

func (l *HashLiteral) TokenLiteral() string {
	return l.Token.Literal
}

func (l *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := make([]string, 0)
	for k, v := range l.Pairs {
		pairs = append(pairs, k.String()+":"+v.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (e *IndexExpression) expressionNode() {}

func (e *IndexExpression) TokenLiteral() string {
	return e.Token.Literal
}

func (e *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString("[")
	out.WriteString(e.Index.String())
	out.WriteString("])")
	return out.String()
}

type MacroLiteral struct {
	Token      token.Token // token.MACRO
	Parameters []*Identifier
	Body       *BlockStatement
}

func (l *MacroLiteral) expressionNode() {}

func (l *MacroLiteral) TokenLiteral() string {
	return l.Token.Literal
}

func (l *MacroLiteral) String() string {
	var out bytes.Buffer
	params := make([]string, 0)
	for _, p := range l.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(l.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(l.Body.String())
	return out.String()
}
