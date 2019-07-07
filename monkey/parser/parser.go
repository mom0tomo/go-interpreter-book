package parser

import (
	"github.com/mom0tomo/go-interpreter-book/monkey/ast"
	"github.com/mom0tomo/go-interpreter-book/monkey/lexer"
	"github.com/mom0tomo/go-interpreter-book/monkey/token"
)

type Parser struct {
	l *lexer.Lexer // 字句解析器インスタンスへのポインタ

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つトークンを読み込む。curTokenとpeekTokenの両方がセットされる。
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
