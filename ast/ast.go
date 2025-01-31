package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	exressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier // hold identifier of binding
	Value Expression // expression tht produces value
}

func (ls *LetStatement) statementNode()		{}
func (ls *LetStatemtnt) TokenLiteral() string	{ return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()		{}
func (i *Identifier) TokenLiteral() string	{ return i.Token.Literal }
