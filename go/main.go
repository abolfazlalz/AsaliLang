package main

import (
	"asalicompiler/project"
	"os"

	"asalicompiler/parsing"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing.NewAsaliLangGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewAsaliLangGrammarParser(stream)
	p.BuildParseTrees = true
	tree := p.Parse()
	v := project.NewVisitor()
	v.Visit(tree)
}
