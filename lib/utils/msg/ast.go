package msg

import (
	"bytes"
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

// 構文解析器が生成する全てのASTのルートノードになる
type Program struct {
	Statements []Statement
}

// インターフェースで定義されている関数の1つ
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// インターフェースで定義されている関数の1つ
// 文字列表示してデバッグしやすいようにする
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type ExpressionStatement struct {
	Token      Token      // 式の最初のトークン
	Expression Expression // 式を保持
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type CmdExpression struct {
	Token Token // '['トークン
	Cmd   Expression
}

func (ie *CmdExpression) expressionNode()      {}
func (ie *CmdExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *CmdExpression) String() string {
	var out bytes.Buffer

	out.WriteString("[")
	out.WriteString(ie.Cmd.String())
	out.WriteString("]")

	return out.String()
}

type TextLiteral struct {
	Token Token
	Value string
}

func (sl *TextLiteral) expressionNode()      {}
func (sl *TextLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *TextLiteral) String() string       { return sl.Token.Literal }
