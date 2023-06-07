package main

import (
	"asalicompiler/project"
	"os"

	"asalicompiler/parsing"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing.NewMyGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewMyGrammarParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Parse()
	v := project.NewVisitor()
	v.Visit(tree)
}
