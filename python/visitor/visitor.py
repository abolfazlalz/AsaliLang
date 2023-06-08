import math
from parsing.AsaliLangGrammarLexer import AsaliLangGrammarLexer
from parsing.AsaliLangGrammarVisitor import AsaliLangGrammarVisitor
from parsing.AsaliLangGrammarParser import AsaliLangGrammarParser
from helpers.methods import Methods


class CustomVisitor(AsaliLangGrammarVisitor):
    def __init__(self) -> None:
        super().__init__()
        self.vars = {}
        self.methods = Methods()

    def visitParse(self, ctx):
        self.visit(ctx.block())

    def visitBlock(self, ctx):
        for stat in ctx.stat():
            self.visit(stat)

    def visitStatBlock(self, ctx):
        if ctx.stat() != None:
            return self.visit(ctx.stat())
        elif ctx.block() != None:
            return self.visit(ctx.block())
        return None

    def visitAssignment(self, ctx):
        value = self.visit(ctx.expr())
        if str(value).isnumeric():
            value = float(value)
        self.vars[ctx.ID().getText()] = value
        return super().visitAssignment(ctx)

    def visitMethodCallExpr(self, ctx):
        return self.visit(ctx.inlineMethodCall())

    def visitPowExpr(self, ctx):
        left = self.visit(ctx.expr(0))
        right = self.visit(ctx.expr(1))
        return math.pow(left, right)

    def visitUnaryMinusExpr(self, ctx):
        return -1 * self.visit(ctx.expr())

    def visitMultiplicationExpr(self, ctx):
        left = self.visit(ctx.expr(0))
        right = self.visit(ctx.expr(1))
        op = ctx.op.type
        if op == AsaliLangGrammarLexer.MULT:
            return left * right
        elif op == AsaliLangGrammarLexer.DIV:
            return left / right
        else:
            return left % right

    def visitAdditiveExpr(self, ctx):
        left = self.visit(ctx.expr(0))
        right = self.visit(ctx.expr(1))
        op = ctx.op.type
        if not isinstance(left, float) or not isinstance(right, float):
            return str(left) + str(right)
        if op == AsaliLangGrammarLexer.PLUS:
            return left + right
        else:
            return left - right

    def visitRelationalExpr(self, ctx):
        left = self.visit(ctx.expr(0))
        right = self.visit(ctx.expr(1))
        op = ctx.op.type
        if op == AsaliLangGrammarLexer.LT:
            return left < right
        elif op == AsaliLangGrammarLexer.LTEQ:
            return left <= right
        elif op == AsaliLangGrammarLexer.GT:
            return left > right
        else:
            return left >= right

    def visitEqualityExpr(self, ctx):
        left = self.visit(ctx.expr(0))
        right = self.visit(ctx.expr(1))
        op = ctx.op.type
        if op == AsaliLangGrammarLexer.EQ:
            return left == right
        elif op == AsaliLangGrammarLexer.NEQ:
            return left != right

    def visitAndExpr(self, ctx):
        left = self.visit(ctx.expr(0))
        right = self.visit(ctx.expr(1))
        return left and right

    def visitOrExpr(self, ctx):
        left = self.visit(ctx.expr(0))
        right = self.visit(ctx.expr(1))
        return left or right

    def visitNotExpr(self, ctx):
        return not self.visit(ctx.expr())

    def visitParExpr(self, ctx):
        return self.visit(ctx.expr())

    def visitAtomExpr(self, ctx):
        return self.visit(ctx.atom())

    def visitIdAtom(self, ctx):
        re = self.vars[ctx.ID().getText()]
        return re

    def visitStringAtom(self, ctx):
        value = ctx.getText()
        value: str = value[1:len(value) - 1]
        return value.replace("\"\"", "\"")

    def visitNumberAtom(self, ctx):
        value = ctx.getText()
        return float(value)

    def visitBooleanAtom(self, ctx):
        return ctx.TRUE() != None

    def visitMethodCall(self, ctx):
        return self.handleMethodCall(ctx)

    def visitInlineMethodCall(self, ctx):
        return self.handleMethodCall(ctx)

    def handleMethodCall(self, ctx):
        args = self.visit(ctx.methodCallArguments())
        methodName = ctx.ID().getText()
        if methodName == "print":
            self.methods.print(args)
        if methodName == "sin":
            return self.methods.sin(args)
        return None

    def visitMethodCallArguments(self, ctx):
        exprList = ctx.expr()
        result = []
        for expr in exprList:
            result.append(self.visit(expr))
        return result

    def visitMethodCallStat(self, ctx):
        return self.visit(ctx.methodCall())

    def visitForStat(self, ctx):
        varName = ctx.ID().getText()
        start = int(self.visit(ctx.expr(0)))
        end = int(self.visit(ctx.expr(1)))

        # Last variable value for this name
        self.vars[varName] = start

        # for i in range(start, end):
        while self.vars[varName] < end:
            self.visit(ctx.statBlock())
            self.vars[varName] = self.vars[varName] + 1

    def visitLoopStat(self, ctx):
        varName = ctx.ID().getText()
        end = int(self.visit(ctx.expr()))

        self.vars[varName] = 1
        while self.vars[varName] <= end:
            self.visit(ctx.statBlock())
            self.vars[varName] = self.vars[varName] + 1

    def visitWhileStat(self, ctx):
        result = self.visit(ctx.expr())
        while result:
            self.visit(ctx.statBlock())
            result = self.visit(ctx.expr())

    def visitIfStat(self, ctx):
        conditions = ctx.conditionBlock()
        evaluatedBlock = False
        for condition in conditions:
            evaluated = self.visit(condition.expr())
            if evaluated:
                evaluatedBlock = True
                self.visit(condition.statBlock())
                break

        if not evaluatedBlock:
            if ctx.statBlock():
                self.visit(ctx.statBlock())
