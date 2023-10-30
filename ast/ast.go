package ast

import (
	"bytes"
	"monkey/token"
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
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token //token.LET トークン
	Name  *Identifier
	Calue Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token //token.LET トークン
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type RetrunStatement struct {
	Token       token.Token // 'return'トークン
	RetrunValue Expression
}

func (rs *RetrunStatement) statementNode()       {}
func (rs *RetrunStatement) TokenLiteral() string { return rs.Token.Literal }

type ExpressionStatement struct {
	Token      token.Token //　式の最初のトークン
	Expression Expression
}

func (es ExpressionStatement) statementNode()       {}
func (es ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
