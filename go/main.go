package main

import (
	"os"

	project "github.com/abolfazlalz/AsaliLang/project"

	parsing "github.com/abolfazlalz/AsaliLang/parsing"

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
