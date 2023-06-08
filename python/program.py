from antlr4 import *

from visitor.visitor import CustomListener
from parsing.AsaliLangGrammarLexer import AsaliLangGrammarLexer
from parsing.AsaliLangGrammarParser import AsaliLangGrammarParser


def main():
    input_stream = InputStream("print hello;")
    lexer = AsaliLangGrammarLexer(input_stream)
    stream = CommonTokenStream(lexer)
    parser = AsaliLangGrammarParser(stream)
    tree = parser.parse()
    printer = CustomListener()
    walker = ParseTreeWalker()
    walker.walk(printer, tree)


if __name__ == '__main__':
    main()
