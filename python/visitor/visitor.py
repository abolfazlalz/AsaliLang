from parsing.AsaliLangGrammarVisitor import AsaliLangGrammarVisitor


class CustomVisitor(AsaliLangGrammarVisitor):
    def visitParse(self, ctx):
        self.visit
        print("hello world")
