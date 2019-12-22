package main

import (
	"decode/src/ast"
	"decode/src/parser"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	args := os.Args[1]
	Parse(args)
}

func Parse(expr string) {
	input := antlr.NewInputStream(expr)
	lexer := parser.NewzipLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parserer := parser.NewzipParser(tokenStream)
	tree := parserer.Unit()
	v := &parser.ZipVisitor{}
	unit := v.Visit(tree).(*ast.Unit)
	println(unit.Eval())
}
