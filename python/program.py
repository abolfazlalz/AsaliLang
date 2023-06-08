import sys

from antlr4 import *

from visitor.visitor import CustomVisitor
from parsing.AsaliLangGrammarLexer import AsaliLangGrammarLexer
from parsing.AsaliLangGrammarParser import AsaliLangGrammarParser


def main():
    if len(sys.argv) < 2:
        print("Error: Enter a file path")
        return
    # Read from file in argv
    with open(sys.argv[1]) as file:
        data = f'{file.read()}\n'

    input_stream = InputStream(data)
    lexer = AsaliLangGrammarLexer(input_stream)
    stream = CommonTokenStream(lexer)
    parser = AsaliLangGrammarParser(stream)
    tree = parser.parse()
    visitor = CustomVisitor()
    visitor.visit(tree)


if __name__ == '__main__':
    main()
