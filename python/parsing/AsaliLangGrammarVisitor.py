# Generated from AsaliLangGrammar.g4 by ANTLR 4.13.0
from antlr4 import *

# This class defines a complete generic visitor for a parse tree produced by AsaliLangGrammarParser.

class AsaliLangGrammarVisitor(ParseTreeVisitor):

    # Visit a parse tree produced by AsaliLangGrammarParser#parse.
    def visitParse(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#block.
    def visitBlock(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#stat.
    def visitStat(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#assignment.
    def visitAssignment(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#ifStat.
    def visitIfStat(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#conditionBlock.
    def visitConditionBlock(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#statBlock.
    def visitStatBlock(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#whileStat.
    def visitWhileStat(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#forStat.
    def visitForStat(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#loopStat.
    def visitLoopStat(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#methodCallStat.
    def visitMethodCallStat(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#methodCall.
    def visitMethodCall(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#inlineMethodCall.
    def visitInlineMethodCall(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#methodCallArguments.
    def visitMethodCallArguments(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#methodCallExpr.
    def visitMethodCallExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#notExpr.
    def visitNotExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#unaryMinusExpr.
    def visitUnaryMinusExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#multiplicationExpr.
    def visitMultiplicationExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#atomExpr.
    def visitAtomExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#orExpr.
    def visitOrExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#additiveExpr.
    def visitAdditiveExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#powExpr.
    def visitPowExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#relationalExpr.
    def visitRelationalExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#equalityExpr.
    def visitEqualityExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#andExpr.
    def visitAndExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#parExpr.
    def visitParExpr(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#numberAtom.
    def visitNumberAtom(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#booleanAtom.
    def visitBooleanAtom(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#idAtom.
    def visitIdAtom(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#stringAtom.
    def visitStringAtom(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AsaliLangGrammarParser#nilAtom.
    def visitNilAtom(self, ctx):
        return self.visitChildren(ctx)


