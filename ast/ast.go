package ast

import "github.com/saraikium/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Statements []Statement
}

// assert LetStatement implementation
var _ Node = (*LetStatement)(nil)

type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) StatementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// assert Identifier implementation
var _ Expression = (*Identifier)(nil)

type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

// Assert ReturnStatement implementation
var _ Statement = (*ReturnStatement)(nil)

type ReturnStatement struct {
	Token token.Token // The 'return' key word token
	Value Expression
}

func (rs *ReturnStatement) StatementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (i *Identifier) ExpressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
