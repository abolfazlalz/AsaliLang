# Generated from ./AsaliLangGrammar.g4 by ANTLR 4.13.0
# encoding: utf-8
from __future__ import print_function
from antlr4 import *
from io import StringIO
import sys

def serializedATN():
    return [
        4,1,42,197,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,
        6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,
        2,14,7,14,2,15,7,15,2,16,7,16,2,17,7,17,1,0,1,0,1,0,1,1,5,1,41,8,
        1,10,1,12,1,44,9,1,1,2,1,2,1,2,1,2,1,2,1,2,1,2,3,2,53,8,2,1,3,1,
        3,1,3,1,3,1,3,1,4,1,4,1,4,1,4,1,4,5,4,65,8,4,10,4,12,4,68,9,4,1,
        4,1,4,3,4,72,8,4,1,5,1,5,1,5,1,6,1,6,1,6,1,6,1,6,3,6,82,8,6,1,6,
        1,6,1,6,1,6,1,6,1,6,1,6,3,6,91,8,6,1,7,1,7,1,7,1,7,1,8,1,8,1,8,1,
        8,1,8,1,8,1,8,1,8,1,9,1,9,1,9,1,9,1,9,1,9,1,10,1,10,1,10,1,11,1,
        11,1,11,1,11,1,11,1,11,1,11,1,12,1,12,1,12,1,13,1,13,1,13,1,13,1,
        13,1,14,1,14,1,14,1,14,5,14,133,8,14,10,14,12,14,136,9,14,3,14,138,
        8,14,1,15,1,15,1,15,1,15,5,15,144,8,15,10,15,12,15,147,9,15,3,15,
        149,8,15,1,16,1,16,1,16,1,16,1,16,1,16,1,16,3,16,158,8,16,1,16,1,
        16,1,16,1,16,1,16,1,16,1,16,1,16,1,16,1,16,1,16,1,16,1,16,1,16,1,
        16,1,16,1,16,1,16,1,16,1,16,1,16,5,16,181,8,16,10,16,12,16,184,9,
        16,1,17,1,17,1,17,1,17,1,17,1,17,1,17,1,17,1,17,3,17,195,8,17,1,
        17,0,1,32,18,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,0,
        6,1,0,12,14,1,0,10,11,1,0,6,9,1,0,4,5,1,0,38,39,1,0,28,29,210,0,
        36,1,0,0,0,2,42,1,0,0,0,4,52,1,0,0,0,6,54,1,0,0,0,8,59,1,0,0,0,10,
        73,1,0,0,0,12,90,1,0,0,0,14,92,1,0,0,0,16,96,1,0,0,0,18,104,1,0,
        0,0,20,110,1,0,0,0,22,113,1,0,0,0,24,120,1,0,0,0,26,123,1,0,0,0,
        28,137,1,0,0,0,30,148,1,0,0,0,32,157,1,0,0,0,34,194,1,0,0,0,36,37,
        3,2,1,0,37,38,5,0,0,1,38,1,1,0,0,0,39,41,3,4,2,0,40,39,1,0,0,0,41,
        44,1,0,0,0,42,40,1,0,0,0,42,43,1,0,0,0,43,3,1,0,0,0,44,42,1,0,0,
        0,45,53,3,6,3,0,46,53,3,8,4,0,47,53,3,14,7,0,48,53,3,20,10,0,49,
        53,3,16,8,0,50,53,3,18,9,0,51,53,3,22,11,0,52,45,1,0,0,0,52,46,1,
        0,0,0,52,47,1,0,0,0,52,48,1,0,0,0,52,49,1,0,0,0,52,50,1,0,0,0,52,
        51,1,0,0,0,53,5,1,0,0,0,54,55,5,37,0,0,55,56,5,19,0,0,56,57,3,32,
        16,0,57,58,5,17,0,0,58,7,1,0,0,0,59,60,5,31,0,0,60,66,3,10,5,0,61,
        62,5,32,0,0,62,63,5,31,0,0,63,65,3,10,5,0,64,61,1,0,0,0,65,68,1,
        0,0,0,66,64,1,0,0,0,66,67,1,0,0,0,67,71,1,0,0,0,68,66,1,0,0,0,69,
        70,5,32,0,0,70,72,3,12,6,0,71,69,1,0,0,0,71,72,1,0,0,0,72,9,1,0,
        0,0,73,74,3,32,16,0,74,75,3,12,6,0,75,11,1,0,0,0,76,77,5,22,0,0,
        77,78,3,2,1,0,78,79,5,23,0,0,79,91,1,0,0,0,80,82,5,26,0,0,81,80,
        1,0,0,0,81,82,1,0,0,0,82,83,1,0,0,0,83,84,5,24,0,0,84,85,3,2,1,0,
        85,86,5,25,0,0,86,91,1,0,0,0,87,91,3,4,2,0,88,89,5,27,0,0,89,91,
        3,2,1,0,90,76,1,0,0,0,90,81,1,0,0,0,90,87,1,0,0,0,90,88,1,0,0,0,
        91,13,1,0,0,0,92,93,5,33,0,0,93,94,3,32,16,0,94,95,3,12,6,0,95,15,
        1,0,0,0,96,97,5,34,0,0,97,98,5,37,0,0,98,99,5,19,0,0,99,100,3,32,
        16,0,100,101,5,18,0,0,101,102,3,32,16,0,102,103,3,12,6,0,103,17,
        1,0,0,0,104,105,5,35,0,0,105,106,5,37,0,0,106,107,5,18,0,0,107,108,
        3,32,16,0,108,109,3,12,6,0,109,19,1,0,0,0,110,111,3,24,12,0,111,
        112,5,17,0,0,112,21,1,0,0,0,113,114,5,36,0,0,114,115,5,37,0,0,115,
        116,5,20,0,0,116,117,3,30,15,0,117,118,5,21,0,0,118,119,3,12,6,0,
        119,23,1,0,0,0,120,121,5,37,0,0,121,122,3,28,14,0,122,25,1,0,0,0,
        123,124,5,37,0,0,124,125,5,20,0,0,125,126,3,28,14,0,126,127,5,21,
        0,0,127,27,1,0,0,0,128,138,1,0,0,0,129,134,3,32,16,0,130,131,5,1,
        0,0,131,133,3,32,16,0,132,130,1,0,0,0,133,136,1,0,0,0,134,132,1,
        0,0,0,134,135,1,0,0,0,135,138,1,0,0,0,136,134,1,0,0,0,137,128,1,
        0,0,0,137,129,1,0,0,0,138,29,1,0,0,0,139,149,1,0,0,0,140,145,5,37,
        0,0,141,142,5,1,0,0,142,144,5,37,0,0,143,141,1,0,0,0,144,147,1,0,
        0,0,145,143,1,0,0,0,145,146,1,0,0,0,146,149,1,0,0,0,147,145,1,0,
        0,0,148,139,1,0,0,0,148,140,1,0,0,0,149,31,1,0,0,0,150,151,6,16,
        -1,0,151,158,3,26,13,0,152,153,5,11,0,0,153,158,3,32,16,9,154,155,
        5,16,0,0,155,158,3,32,16,8,156,158,3,34,17,0,157,150,1,0,0,0,157,
        152,1,0,0,0,157,154,1,0,0,0,157,156,1,0,0,0,158,182,1,0,0,0,159,
        160,10,10,0,0,160,161,5,15,0,0,161,181,3,32,16,10,162,163,10,7,0,
        0,163,164,7,0,0,0,164,181,3,32,16,8,165,166,10,6,0,0,166,167,7,1,
        0,0,167,181,3,32,16,7,168,169,10,5,0,0,169,170,7,2,0,0,170,181,3,
        32,16,6,171,172,10,4,0,0,172,173,7,3,0,0,173,181,3,32,16,5,174,175,
        10,3,0,0,175,176,5,3,0,0,176,181,3,32,16,4,177,178,10,2,0,0,178,
        179,5,2,0,0,179,181,3,32,16,3,180,159,1,0,0,0,180,162,1,0,0,0,180,
        165,1,0,0,0,180,168,1,0,0,0,180,171,1,0,0,0,180,174,1,0,0,0,180,
        177,1,0,0,0,181,184,1,0,0,0,182,180,1,0,0,0,182,183,1,0,0,0,183,
        33,1,0,0,0,184,182,1,0,0,0,185,186,5,20,0,0,186,187,3,32,16,0,187,
        188,5,21,0,0,188,195,1,0,0,0,189,195,7,4,0,0,190,195,7,5,0,0,191,
        195,5,37,0,0,192,195,5,40,0,0,193,195,5,30,0,0,194,185,1,0,0,0,194,
        189,1,0,0,0,194,190,1,0,0,0,194,191,1,0,0,0,194,192,1,0,0,0,194,
        193,1,0,0,0,195,35,1,0,0,0,14,42,52,66,71,81,90,134,137,145,148,
        157,180,182,194
    ]

class AsaliLangGrammarParser ( Parser ):

    grammarFileName = "AsaliLangGrammar.g4"

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    sharedContextCache = PredictionContextCache()

    literalNames = [ u"<INVALID>", u"','", u"'||'", u"'&&'", u"'=='", u"'!='", 
                     u"'>'", u"'<'", u"'>='", u"'<='", u"'+'", u"'-'", u"'*'", 
                     u"'/'", u"'%'", u"'^'", u"'!'", u"';'", u"':'", u"'='", 
                     u"'('", u"')'", u"'{'", u"'}'", u"'begin'", u"'end'", 
                     u"'do'", u"'then'", u"'true'", u"'false'", u"'nil'", 
                     u"'if'", u"'else'", u"'while'", u"'for'", u"'loop'", 
                     u"'func'" ]

    symbolicNames = [ u"<INVALID>", u"<INVALID>", u"OR", u"AND", u"EQ", 
                      u"NEQ", u"GT", u"LT", u"GTEQ", u"LTEQ", u"PLUS", u"MINUS", 
                      u"MULT", u"DIV", u"MOD", u"POW", u"NOT", u"SCOL", 
                      u"COL", u"ASSIGN", u"OPAR", u"CPAR", u"OBRACE", u"CBRACE", 
                      u"BEGIN", u"END", u"DO", u"THEN", u"TRUE", u"FALSE", 
                      u"NIL", u"IF", u"ELSE", u"WHILE", u"FOR", u"LOOP", 
                      u"FUNC", u"ID", u"INT", u"FLOAT", u"STRING", u"COMMENT", 
                      u"SPACE" ]

    RULE_parse = 0
    RULE_block = 1
    RULE_stat = 2
    RULE_assignment = 3
    RULE_ifStat = 4
    RULE_conditionBlock = 5
    RULE_statBlock = 6
    RULE_whileStat = 7
    RULE_forStat = 8
    RULE_loopStat = 9
    RULE_methodCallStat = 10
    RULE_defineFuncStats = 11
    RULE_methodCall = 12
    RULE_inlineMethodCall = 13
    RULE_methodCallArguments = 14
    RULE_defineFuncArguments = 15
    RULE_expr = 16
    RULE_atom = 17

    ruleNames =  [ u"parse", u"block", u"stat", u"assignment", u"ifStat", 
                   u"conditionBlock", u"statBlock", u"whileStat", u"forStat", 
                   u"loopStat", u"methodCallStat", u"defineFuncStats", u"methodCall", 
                   u"inlineMethodCall", u"methodCallArguments", u"defineFuncArguments", 
                   u"expr", u"atom" ]

    EOF = Token.EOF
    T__0=1
    OR=2
    AND=3
    EQ=4
    NEQ=5
    GT=6
    LT=7
    GTEQ=8
    LTEQ=9
    PLUS=10
    MINUS=11
    MULT=12
    DIV=13
    MOD=14
    POW=15
    NOT=16
    SCOL=17
    COL=18
    ASSIGN=19
    OPAR=20
    CPAR=21
    OBRACE=22
    CBRACE=23
    BEGIN=24
    END=25
    DO=26
    THEN=27
    TRUE=28
    FALSE=29
    NIL=30
    IF=31
    ELSE=32
    WHILE=33
    FOR=34
    LOOP=35
    FUNC=36
    ID=37
    INT=38
    FLOAT=39
    STRING=40
    COMMENT=41
    SPACE=42

    def __init__(self, input, output=sys.stdout):
        super(AsaliLangGrammarParser, self).__init__(input, output=output)
        self.checkVersion("4.13.0")
        self._interp = ParserATNSimulator(self, self.atn, self.decisionsToDFA, self.sharedContextCache)
        self._predicates = None




    class ParseContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.ParseContext, self).__init__(parent, invokingState)
            self.parser = parser

        def block(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.BlockContext,0)


        def EOF(self):
            return self.getToken(AsaliLangGrammarParser.EOF, 0)

        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_parse

        def enterRule(self, listener):
            if hasattr(listener, "enterParse"):
                listener.enterParse(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitParse"):
                listener.exitParse(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitParse"):
                return visitor.visitParse(self)
            else:
                return visitor.visitChildren(self)




    def parse(self):

        localctx = AsaliLangGrammarParser.ParseContext(self, self._ctx, self.state)
        self.enterRule(localctx, 0, self.RULE_parse)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 36
            self.block()
            self.state = 37
            self.match(AsaliLangGrammarParser.EOF)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class BlockContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.BlockContext, self).__init__(parent, invokingState)
            self.parser = parser

        def stat(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.StatContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.StatContext,i)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_block

        def enterRule(self, listener):
            if hasattr(listener, "enterBlock"):
                listener.enterBlock(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitBlock"):
                listener.exitBlock(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitBlock"):
                return visitor.visitBlock(self)
            else:
                return visitor.visitChildren(self)




    def block(self):

        localctx = AsaliLangGrammarParser.BlockContext(self, self._ctx, self.state)
        self.enterRule(localctx, 2, self.RULE_block)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 42
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,0,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    self.state = 39
                    self.stat() 
                self.state = 44
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,0,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class StatContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.StatContext, self).__init__(parent, invokingState)
            self.parser = parser

        def assignment(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.AssignmentContext,0)


        def ifStat(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.IfStatContext,0)


        def whileStat(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.WhileStatContext,0)


        def methodCallStat(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.MethodCallStatContext,0)


        def forStat(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.ForStatContext,0)


        def loopStat(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.LoopStatContext,0)


        def defineFuncStats(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.DefineFuncStatsContext,0)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_stat

        def enterRule(self, listener):
            if hasattr(listener, "enterStat"):
                listener.enterStat(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitStat"):
                listener.exitStat(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitStat"):
                return visitor.visitStat(self)
            else:
                return visitor.visitChildren(self)




    def stat(self):

        localctx = AsaliLangGrammarParser.StatContext(self, self._ctx, self.state)
        self.enterRule(localctx, 4, self.RULE_stat)
        try:
            self.state = 52
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,1,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 45
                self.assignment()
                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 46
                self.ifStat()
                pass

            elif la_ == 3:
                self.enterOuterAlt(localctx, 3)
                self.state = 47
                self.whileStat()
                pass

            elif la_ == 4:
                self.enterOuterAlt(localctx, 4)
                self.state = 48
                self.methodCallStat()
                pass

            elif la_ == 5:
                self.enterOuterAlt(localctx, 5)
                self.state = 49
                self.forStat()
                pass

            elif la_ == 6:
                self.enterOuterAlt(localctx, 6)
                self.state = 50
                self.loopStat()
                pass

            elif la_ == 7:
                self.enterOuterAlt(localctx, 7)
                self.state = 51
                self.defineFuncStats()
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class AssignmentContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.AssignmentContext, self).__init__(parent, invokingState)
            self.parser = parser

        def ID(self):
            return self.getToken(AsaliLangGrammarParser.ID, 0)

        def ASSIGN(self):
            return self.getToken(AsaliLangGrammarParser.ASSIGN, 0)

        def expr(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,0)


        def SCOL(self):
            return self.getToken(AsaliLangGrammarParser.SCOL, 0)

        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_assignment

        def enterRule(self, listener):
            if hasattr(listener, "enterAssignment"):
                listener.enterAssignment(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitAssignment"):
                listener.exitAssignment(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitAssignment"):
                return visitor.visitAssignment(self)
            else:
                return visitor.visitChildren(self)




    def assignment(self):

        localctx = AsaliLangGrammarParser.AssignmentContext(self, self._ctx, self.state)
        self.enterRule(localctx, 6, self.RULE_assignment)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 54
            self.match(AsaliLangGrammarParser.ID)
            self.state = 55
            self.match(AsaliLangGrammarParser.ASSIGN)
            self.state = 56
            self.expr(0)
            self.state = 57
            self.match(AsaliLangGrammarParser.SCOL)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class IfStatContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.IfStatContext, self).__init__(parent, invokingState)
            self.parser = parser

        def IF(self, i=None):
            if i is None:
                return self.getTokens(AsaliLangGrammarParser.IF)
            else:
                return self.getToken(AsaliLangGrammarParser.IF, i)

        def conditionBlock(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ConditionBlockContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ConditionBlockContext,i)


        def ELSE(self, i=None):
            if i is None:
                return self.getTokens(AsaliLangGrammarParser.ELSE)
            else:
                return self.getToken(AsaliLangGrammarParser.ELSE, i)

        def statBlock(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.StatBlockContext,0)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_ifStat

        def enterRule(self, listener):
            if hasattr(listener, "enterIfStat"):
                listener.enterIfStat(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitIfStat"):
                listener.exitIfStat(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitIfStat"):
                return visitor.visitIfStat(self)
            else:
                return visitor.visitChildren(self)




    def ifStat(self):

        localctx = AsaliLangGrammarParser.IfStatContext(self, self._ctx, self.state)
        self.enterRule(localctx, 8, self.RULE_ifStat)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 59
            self.match(AsaliLangGrammarParser.IF)
            self.state = 60
            self.conditionBlock()
            self.state = 66
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,2,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    self.state = 61
                    self.match(AsaliLangGrammarParser.ELSE)
                    self.state = 62
                    self.match(AsaliLangGrammarParser.IF)
                    self.state = 63
                    self.conditionBlock() 
                self.state = 68
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,2,self._ctx)

            self.state = 71
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,3,self._ctx)
            if la_ == 1:
                self.state = 69
                self.match(AsaliLangGrammarParser.ELSE)
                self.state = 70
                self.statBlock()


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ConditionBlockContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.ConditionBlockContext, self).__init__(parent, invokingState)
            self.parser = parser

        def expr(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,0)


        def statBlock(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.StatBlockContext,0)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_conditionBlock

        def enterRule(self, listener):
            if hasattr(listener, "enterConditionBlock"):
                listener.enterConditionBlock(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitConditionBlock"):
                listener.exitConditionBlock(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitConditionBlock"):
                return visitor.visitConditionBlock(self)
            else:
                return visitor.visitChildren(self)




    def conditionBlock(self):

        localctx = AsaliLangGrammarParser.ConditionBlockContext(self, self._ctx, self.state)
        self.enterRule(localctx, 10, self.RULE_conditionBlock)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 73
            self.expr(0)
            self.state = 74
            self.statBlock()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class StatBlockContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.StatBlockContext, self).__init__(parent, invokingState)
            self.parser = parser

        def OBRACE(self):
            return self.getToken(AsaliLangGrammarParser.OBRACE, 0)

        def block(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.BlockContext,0)


        def CBRACE(self):
            return self.getToken(AsaliLangGrammarParser.CBRACE, 0)

        def BEGIN(self):
            return self.getToken(AsaliLangGrammarParser.BEGIN, 0)

        def END(self):
            return self.getToken(AsaliLangGrammarParser.END, 0)

        def DO(self):
            return self.getToken(AsaliLangGrammarParser.DO, 0)

        def stat(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.StatContext,0)


        def THEN(self):
            return self.getToken(AsaliLangGrammarParser.THEN, 0)

        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_statBlock

        def enterRule(self, listener):
            if hasattr(listener, "enterStatBlock"):
                listener.enterStatBlock(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitStatBlock"):
                listener.exitStatBlock(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitStatBlock"):
                return visitor.visitStatBlock(self)
            else:
                return visitor.visitChildren(self)




    def statBlock(self):

        localctx = AsaliLangGrammarParser.StatBlockContext(self, self._ctx, self.state)
        self.enterRule(localctx, 12, self.RULE_statBlock)
        self._la = 0 # Token type
        try:
            self.state = 90
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [22]:
                self.enterOuterAlt(localctx, 1)
                self.state = 76
                self.match(AsaliLangGrammarParser.OBRACE)
                self.state = 77
                self.block()
                self.state = 78
                self.match(AsaliLangGrammarParser.CBRACE)
                pass
            elif token in [24, 26]:
                self.enterOuterAlt(localctx, 2)
                self.state = 81
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if _la==26:
                    self.state = 80
                    self.match(AsaliLangGrammarParser.DO)


                self.state = 83
                self.match(AsaliLangGrammarParser.BEGIN)
                self.state = 84
                self.block()
                self.state = 85
                self.match(AsaliLangGrammarParser.END)
                pass
            elif token in [31, 33, 34, 35, 36, 37]:
                self.enterOuterAlt(localctx, 3)
                self.state = 87
                self.stat()
                pass
            elif token in [27]:
                self.enterOuterAlt(localctx, 4)
                self.state = 88
                self.match(AsaliLangGrammarParser.THEN)
                self.state = 89
                self.block()
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class WhileStatContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.WhileStatContext, self).__init__(parent, invokingState)
            self.parser = parser

        def WHILE(self):
            return self.getToken(AsaliLangGrammarParser.WHILE, 0)

        def expr(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,0)


        def statBlock(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.StatBlockContext,0)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_whileStat

        def enterRule(self, listener):
            if hasattr(listener, "enterWhileStat"):
                listener.enterWhileStat(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitWhileStat"):
                listener.exitWhileStat(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitWhileStat"):
                return visitor.visitWhileStat(self)
            else:
                return visitor.visitChildren(self)




    def whileStat(self):

        localctx = AsaliLangGrammarParser.WhileStatContext(self, self._ctx, self.state)
        self.enterRule(localctx, 14, self.RULE_whileStat)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 92
            self.match(AsaliLangGrammarParser.WHILE)
            self.state = 93
            self.expr(0)
            self.state = 94
            self.statBlock()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ForStatContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.ForStatContext, self).__init__(parent, invokingState)
            self.parser = parser

        def FOR(self):
            return self.getToken(AsaliLangGrammarParser.FOR, 0)

        def ID(self):
            return self.getToken(AsaliLangGrammarParser.ID, 0)

        def ASSIGN(self):
            return self.getToken(AsaliLangGrammarParser.ASSIGN, 0)

        def expr(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ExprContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,i)


        def COL(self):
            return self.getToken(AsaliLangGrammarParser.COL, 0)

        def statBlock(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.StatBlockContext,0)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_forStat

        def enterRule(self, listener):
            if hasattr(listener, "enterForStat"):
                listener.enterForStat(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitForStat"):
                listener.exitForStat(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitForStat"):
                return visitor.visitForStat(self)
            else:
                return visitor.visitChildren(self)




    def forStat(self):

        localctx = AsaliLangGrammarParser.ForStatContext(self, self._ctx, self.state)
        self.enterRule(localctx, 16, self.RULE_forStat)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 96
            self.match(AsaliLangGrammarParser.FOR)
            self.state = 97
            self.match(AsaliLangGrammarParser.ID)
            self.state = 98
            self.match(AsaliLangGrammarParser.ASSIGN)
            self.state = 99
            self.expr(0)
            self.state = 100
            self.match(AsaliLangGrammarParser.COL)
            self.state = 101
            self.expr(0)
            self.state = 102
            self.statBlock()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class LoopStatContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.LoopStatContext, self).__init__(parent, invokingState)
            self.parser = parser

        def LOOP(self):
            return self.getToken(AsaliLangGrammarParser.LOOP, 0)

        def ID(self):
            return self.getToken(AsaliLangGrammarParser.ID, 0)

        def COL(self):
            return self.getToken(AsaliLangGrammarParser.COL, 0)

        def expr(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,0)


        def statBlock(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.StatBlockContext,0)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_loopStat

        def enterRule(self, listener):
            if hasattr(listener, "enterLoopStat"):
                listener.enterLoopStat(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitLoopStat"):
                listener.exitLoopStat(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitLoopStat"):
                return visitor.visitLoopStat(self)
            else:
                return visitor.visitChildren(self)




    def loopStat(self):

        localctx = AsaliLangGrammarParser.LoopStatContext(self, self._ctx, self.state)
        self.enterRule(localctx, 18, self.RULE_loopStat)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 104
            self.match(AsaliLangGrammarParser.LOOP)
            self.state = 105
            self.match(AsaliLangGrammarParser.ID)
            self.state = 106
            self.match(AsaliLangGrammarParser.COL)
            self.state = 107
            self.expr(0)
            self.state = 108
            self.statBlock()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class MethodCallStatContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.MethodCallStatContext, self).__init__(parent, invokingState)
            self.parser = parser

        def methodCall(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.MethodCallContext,0)


        def SCOL(self):
            return self.getToken(AsaliLangGrammarParser.SCOL, 0)

        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_methodCallStat

        def enterRule(self, listener):
            if hasattr(listener, "enterMethodCallStat"):
                listener.enterMethodCallStat(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitMethodCallStat"):
                listener.exitMethodCallStat(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitMethodCallStat"):
                return visitor.visitMethodCallStat(self)
            else:
                return visitor.visitChildren(self)




    def methodCallStat(self):

        localctx = AsaliLangGrammarParser.MethodCallStatContext(self, self._ctx, self.state)
        self.enterRule(localctx, 20, self.RULE_methodCallStat)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 110
            self.methodCall()
            self.state = 111
            self.match(AsaliLangGrammarParser.SCOL)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class DefineFuncStatsContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.DefineFuncStatsContext, self).__init__(parent, invokingState)
            self.parser = parser

        def FUNC(self):
            return self.getToken(AsaliLangGrammarParser.FUNC, 0)

        def ID(self):
            return self.getToken(AsaliLangGrammarParser.ID, 0)

        def OPAR(self):
            return self.getToken(AsaliLangGrammarParser.OPAR, 0)

        def defineFuncArguments(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.DefineFuncArgumentsContext,0)


        def CPAR(self):
            return self.getToken(AsaliLangGrammarParser.CPAR, 0)

        def statBlock(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.StatBlockContext,0)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_defineFuncStats

        def enterRule(self, listener):
            if hasattr(listener, "enterDefineFuncStats"):
                listener.enterDefineFuncStats(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitDefineFuncStats"):
                listener.exitDefineFuncStats(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitDefineFuncStats"):
                return visitor.visitDefineFuncStats(self)
            else:
                return visitor.visitChildren(self)




    def defineFuncStats(self):

        localctx = AsaliLangGrammarParser.DefineFuncStatsContext(self, self._ctx, self.state)
        self.enterRule(localctx, 22, self.RULE_defineFuncStats)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 113
            self.match(AsaliLangGrammarParser.FUNC)
            self.state = 114
            self.match(AsaliLangGrammarParser.ID)
            self.state = 115
            self.match(AsaliLangGrammarParser.OPAR)
            self.state = 116
            self.defineFuncArguments()
            self.state = 117
            self.match(AsaliLangGrammarParser.CPAR)
            self.state = 118
            self.statBlock()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class MethodCallContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.MethodCallContext, self).__init__(parent, invokingState)
            self.parser = parser

        def ID(self):
            return self.getToken(AsaliLangGrammarParser.ID, 0)

        def methodCallArguments(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.MethodCallArgumentsContext,0)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_methodCall

        def enterRule(self, listener):
            if hasattr(listener, "enterMethodCall"):
                listener.enterMethodCall(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitMethodCall"):
                listener.exitMethodCall(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitMethodCall"):
                return visitor.visitMethodCall(self)
            else:
                return visitor.visitChildren(self)




    def methodCall(self):

        localctx = AsaliLangGrammarParser.MethodCallContext(self, self._ctx, self.state)
        self.enterRule(localctx, 24, self.RULE_methodCall)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 120
            self.match(AsaliLangGrammarParser.ID)
            self.state = 121
            self.methodCallArguments()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class InlineMethodCallContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.InlineMethodCallContext, self).__init__(parent, invokingState)
            self.parser = parser

        def ID(self):
            return self.getToken(AsaliLangGrammarParser.ID, 0)

        def OPAR(self):
            return self.getToken(AsaliLangGrammarParser.OPAR, 0)

        def methodCallArguments(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.MethodCallArgumentsContext,0)


        def CPAR(self):
            return self.getToken(AsaliLangGrammarParser.CPAR, 0)

        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_inlineMethodCall

        def enterRule(self, listener):
            if hasattr(listener, "enterInlineMethodCall"):
                listener.enterInlineMethodCall(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitInlineMethodCall"):
                listener.exitInlineMethodCall(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitInlineMethodCall"):
                return visitor.visitInlineMethodCall(self)
            else:
                return visitor.visitChildren(self)




    def inlineMethodCall(self):

        localctx = AsaliLangGrammarParser.InlineMethodCallContext(self, self._ctx, self.state)
        self.enterRule(localctx, 26, self.RULE_inlineMethodCall)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 123
            self.match(AsaliLangGrammarParser.ID)
            self.state = 124
            self.match(AsaliLangGrammarParser.OPAR)
            self.state = 125
            self.methodCallArguments()
            self.state = 126
            self.match(AsaliLangGrammarParser.CPAR)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class MethodCallArgumentsContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.MethodCallArgumentsContext, self).__init__(parent, invokingState)
            self.parser = parser

        def expr(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ExprContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,i)


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_methodCallArguments

        def enterRule(self, listener):
            if hasattr(listener, "enterMethodCallArguments"):
                listener.enterMethodCallArguments(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitMethodCallArguments"):
                listener.exitMethodCallArguments(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitMethodCallArguments"):
                return visitor.visitMethodCallArguments(self)
            else:
                return visitor.visitChildren(self)




    def methodCallArguments(self):

        localctx = AsaliLangGrammarParser.MethodCallArgumentsContext(self, self._ctx, self.state)
        self.enterRule(localctx, 28, self.RULE_methodCallArguments)
        self._la = 0 # Token type
        try:
            self.state = 137
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [17, 21]:
                self.enterOuterAlt(localctx, 1)

                pass
            elif token in [11, 16, 20, 28, 29, 30, 37, 38, 39, 40]:
                self.enterOuterAlt(localctx, 2)
                self.state = 129
                self.expr(0)
                self.state = 134
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while _la==1:
                    self.state = 130
                    self.match(AsaliLangGrammarParser.T__0)
                    self.state = 131
                    self.expr(0)
                    self.state = 136
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)

                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class DefineFuncArgumentsContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.DefineFuncArgumentsContext, self).__init__(parent, invokingState)
            self.parser = parser

        def ID(self, i=None):
            if i is None:
                return self.getTokens(AsaliLangGrammarParser.ID)
            else:
                return self.getToken(AsaliLangGrammarParser.ID, i)

        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_defineFuncArguments

        def enterRule(self, listener):
            if hasattr(listener, "enterDefineFuncArguments"):
                listener.enterDefineFuncArguments(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitDefineFuncArguments"):
                listener.exitDefineFuncArguments(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitDefineFuncArguments"):
                return visitor.visitDefineFuncArguments(self)
            else:
                return visitor.visitChildren(self)




    def defineFuncArguments(self):

        localctx = AsaliLangGrammarParser.DefineFuncArgumentsContext(self, self._ctx, self.state)
        self.enterRule(localctx, 30, self.RULE_defineFuncArguments)
        self._la = 0 # Token type
        try:
            self.state = 148
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [21]:
                self.enterOuterAlt(localctx, 1)

                pass
            elif token in [37]:
                self.enterOuterAlt(localctx, 2)
                self.state = 140
                self.match(AsaliLangGrammarParser.ID)
                self.state = 145
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while _la==1:
                    self.state = 141
                    self.match(AsaliLangGrammarParser.T__0)
                    self.state = 142
                    self.match(AsaliLangGrammarParser.ID)
                    self.state = 147
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)

                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ExprContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.ExprContext, self).__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_expr

     
        def copyFrom(self, ctx):
            super(AsaliLangGrammarParser.ExprContext, self).copyFrom(ctx)


    class MethodCallExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.MethodCallExprContext, self).__init__(parser)
            self.copyFrom(ctx)

        def inlineMethodCall(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.InlineMethodCallContext,0)


        def enterRule(self, listener):
            if hasattr(listener, "enterMethodCallExpr"):
                listener.enterMethodCallExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitMethodCallExpr"):
                listener.exitMethodCallExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitMethodCallExpr"):
                return visitor.visitMethodCallExpr(self)
            else:
                return visitor.visitChildren(self)


    class NotExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.NotExprContext, self).__init__(parser)
            self.copyFrom(ctx)

        def NOT(self):
            return self.getToken(AsaliLangGrammarParser.NOT, 0)
        def expr(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,0)


        def enterRule(self, listener):
            if hasattr(listener, "enterNotExpr"):
                listener.enterNotExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitNotExpr"):
                listener.exitNotExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitNotExpr"):
                return visitor.visitNotExpr(self)
            else:
                return visitor.visitChildren(self)


    class UnaryMinusExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.UnaryMinusExprContext, self).__init__(parser)
            self.copyFrom(ctx)

        def MINUS(self):
            return self.getToken(AsaliLangGrammarParser.MINUS, 0)
        def expr(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,0)


        def enterRule(self, listener):
            if hasattr(listener, "enterUnaryMinusExpr"):
                listener.enterUnaryMinusExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitUnaryMinusExpr"):
                listener.exitUnaryMinusExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitUnaryMinusExpr"):
                return visitor.visitUnaryMinusExpr(self)
            else:
                return visitor.visitChildren(self)


    class MultiplicationExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.MultiplicationExprContext, self).__init__(parser)
            self.op = None # Token
            self.copyFrom(ctx)

        def expr(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ExprContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,i)

        def MULT(self):
            return self.getToken(AsaliLangGrammarParser.MULT, 0)
        def DIV(self):
            return self.getToken(AsaliLangGrammarParser.DIV, 0)
        def MOD(self):
            return self.getToken(AsaliLangGrammarParser.MOD, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterMultiplicationExpr"):
                listener.enterMultiplicationExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitMultiplicationExpr"):
                listener.exitMultiplicationExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitMultiplicationExpr"):
                return visitor.visitMultiplicationExpr(self)
            else:
                return visitor.visitChildren(self)


    class AtomExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.AtomExprContext, self).__init__(parser)
            self.copyFrom(ctx)

        def atom(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.AtomContext,0)


        def enterRule(self, listener):
            if hasattr(listener, "enterAtomExpr"):
                listener.enterAtomExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitAtomExpr"):
                listener.exitAtomExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitAtomExpr"):
                return visitor.visitAtomExpr(self)
            else:
                return visitor.visitChildren(self)


    class OrExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.OrExprContext, self).__init__(parser)
            self.copyFrom(ctx)

        def expr(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ExprContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,i)

        def OR(self):
            return self.getToken(AsaliLangGrammarParser.OR, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterOrExpr"):
                listener.enterOrExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitOrExpr"):
                listener.exitOrExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitOrExpr"):
                return visitor.visitOrExpr(self)
            else:
                return visitor.visitChildren(self)


    class AdditiveExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.AdditiveExprContext, self).__init__(parser)
            self.op = None # Token
            self.copyFrom(ctx)

        def expr(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ExprContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,i)

        def PLUS(self):
            return self.getToken(AsaliLangGrammarParser.PLUS, 0)
        def MINUS(self):
            return self.getToken(AsaliLangGrammarParser.MINUS, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterAdditiveExpr"):
                listener.enterAdditiveExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitAdditiveExpr"):
                listener.exitAdditiveExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitAdditiveExpr"):
                return visitor.visitAdditiveExpr(self)
            else:
                return visitor.visitChildren(self)


    class PowExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.PowExprContext, self).__init__(parser)
            self.copyFrom(ctx)

        def expr(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ExprContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,i)

        def POW(self):
            return self.getToken(AsaliLangGrammarParser.POW, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterPowExpr"):
                listener.enterPowExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitPowExpr"):
                listener.exitPowExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitPowExpr"):
                return visitor.visitPowExpr(self)
            else:
                return visitor.visitChildren(self)


    class RelationalExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.RelationalExprContext, self).__init__(parser)
            self.op = None # Token
            self.copyFrom(ctx)

        def expr(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ExprContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,i)

        def LTEQ(self):
            return self.getToken(AsaliLangGrammarParser.LTEQ, 0)
        def GTEQ(self):
            return self.getToken(AsaliLangGrammarParser.GTEQ, 0)
        def LT(self):
            return self.getToken(AsaliLangGrammarParser.LT, 0)
        def GT(self):
            return self.getToken(AsaliLangGrammarParser.GT, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterRelationalExpr"):
                listener.enterRelationalExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitRelationalExpr"):
                listener.exitRelationalExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitRelationalExpr"):
                return visitor.visitRelationalExpr(self)
            else:
                return visitor.visitChildren(self)


    class EqualityExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.EqualityExprContext, self).__init__(parser)
            self.op = None # Token
            self.copyFrom(ctx)

        def expr(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ExprContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,i)

        def EQ(self):
            return self.getToken(AsaliLangGrammarParser.EQ, 0)
        def NEQ(self):
            return self.getToken(AsaliLangGrammarParser.NEQ, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterEqualityExpr"):
                listener.enterEqualityExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitEqualityExpr"):
                listener.exitEqualityExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitEqualityExpr"):
                return visitor.visitEqualityExpr(self)
            else:
                return visitor.visitChildren(self)


    class AndExprContext(ExprContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.ExprContext)
            super(AsaliLangGrammarParser.AndExprContext, self).__init__(parser)
            self.copyFrom(ctx)

        def expr(self, i=None):
            if i is None:
                return self.getTypedRuleContexts(AsaliLangGrammarParser.ExprContext)
            else:
                return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,i)

        def AND(self):
            return self.getToken(AsaliLangGrammarParser.AND, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterAndExpr"):
                listener.enterAndExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitAndExpr"):
                listener.exitAndExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitAndExpr"):
                return visitor.visitAndExpr(self)
            else:
                return visitor.visitChildren(self)



    def expr(self, _p=0):
        _parentctx = self._ctx
        _parentState = self.state
        localctx = AsaliLangGrammarParser.ExprContext(self, self._ctx, _parentState)
        _prevctx = localctx
        _startState = 32
        self.enterRecursionRule(localctx, 32, self.RULE_expr, _p)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 157
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,10,self._ctx)
            if la_ == 1:
                localctx = AsaliLangGrammarParser.MethodCallExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx

                self.state = 151
                self.inlineMethodCall()
                pass

            elif la_ == 2:
                localctx = AsaliLangGrammarParser.UnaryMinusExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 152
                self.match(AsaliLangGrammarParser.MINUS)
                self.state = 153
                self.expr(9)
                pass

            elif la_ == 3:
                localctx = AsaliLangGrammarParser.NotExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 154
                self.match(AsaliLangGrammarParser.NOT)
                self.state = 155
                self.expr(8)
                pass

            elif la_ == 4:
                localctx = AsaliLangGrammarParser.AtomExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 156
                self.atom()
                pass


            self._ctx.stop = self._input.LT(-1)
            self.state = 182
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,12,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    self.state = 180
                    self._errHandler.sync(self)
                    la_ = self._interp.adaptivePredict(self._input,11,self._ctx)
                    if la_ == 1:
                        localctx = AsaliLangGrammarParser.PowExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 159
                        if not self.precpred(self._ctx, 10):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 10)")
                        self.state = 160
                        self.match(AsaliLangGrammarParser.POW)
                        self.state = 161
                        self.expr(10)
                        pass

                    elif la_ == 2:
                        localctx = AsaliLangGrammarParser.MultiplicationExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 162
                        if not self.precpred(self._ctx, 7):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 7)")
                        self.state = 163
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 28672) != 0)):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 164
                        self.expr(8)
                        pass

                    elif la_ == 3:
                        localctx = AsaliLangGrammarParser.AdditiveExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 165
                        if not self.precpred(self._ctx, 6):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 6)")
                        self.state = 166
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not(_la==10 or _la==11):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 167
                        self.expr(7)
                        pass

                    elif la_ == 4:
                        localctx = AsaliLangGrammarParser.RelationalExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 168
                        if not self.precpred(self._ctx, 5):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 5)")
                        self.state = 169
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 960) != 0)):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 170
                        self.expr(6)
                        pass

                    elif la_ == 5:
                        localctx = AsaliLangGrammarParser.EqualityExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 171
                        if not self.precpred(self._ctx, 4):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 4)")
                        self.state = 172
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not(_la==4 or _la==5):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 173
                        self.expr(5)
                        pass

                    elif la_ == 6:
                        localctx = AsaliLangGrammarParser.AndExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 174
                        if not self.precpred(self._ctx, 3):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 3)")
                        self.state = 175
                        self.match(AsaliLangGrammarParser.AND)
                        self.state = 176
                        self.expr(4)
                        pass

                    elif la_ == 7:
                        localctx = AsaliLangGrammarParser.OrExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 177
                        if not self.precpred(self._ctx, 2):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 2)")
                        self.state = 178
                        self.match(AsaliLangGrammarParser.OR)
                        self.state = 179
                        self.expr(3)
                        pass

             
                self.state = 184
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,12,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.unrollRecursionContexts(_parentctx)
        return localctx


    class AtomContext(ParserRuleContext):

        def __init__(self, parser, parent=None, invokingState=-1):
            super(AsaliLangGrammarParser.AtomContext, self).__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return AsaliLangGrammarParser.RULE_atom

     
        def copyFrom(self, ctx):
            super(AsaliLangGrammarParser.AtomContext, self).copyFrom(ctx)



    class ParExprContext(AtomContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.AtomContext)
            super(AsaliLangGrammarParser.ParExprContext, self).__init__(parser)
            self.copyFrom(ctx)

        def OPAR(self):
            return self.getToken(AsaliLangGrammarParser.OPAR, 0)
        def expr(self):
            return self.getTypedRuleContext(AsaliLangGrammarParser.ExprContext,0)

        def CPAR(self):
            return self.getToken(AsaliLangGrammarParser.CPAR, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterParExpr"):
                listener.enterParExpr(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitParExpr"):
                listener.exitParExpr(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitParExpr"):
                return visitor.visitParExpr(self)
            else:
                return visitor.visitChildren(self)


    class BooleanAtomContext(AtomContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.AtomContext)
            super(AsaliLangGrammarParser.BooleanAtomContext, self).__init__(parser)
            self.copyFrom(ctx)

        def TRUE(self):
            return self.getToken(AsaliLangGrammarParser.TRUE, 0)
        def FALSE(self):
            return self.getToken(AsaliLangGrammarParser.FALSE, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterBooleanAtom"):
                listener.enterBooleanAtom(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitBooleanAtom"):
                listener.exitBooleanAtom(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitBooleanAtom"):
                return visitor.visitBooleanAtom(self)
            else:
                return visitor.visitChildren(self)


    class IdAtomContext(AtomContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.AtomContext)
            super(AsaliLangGrammarParser.IdAtomContext, self).__init__(parser)
            self.copyFrom(ctx)

        def ID(self):
            return self.getToken(AsaliLangGrammarParser.ID, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterIdAtom"):
                listener.enterIdAtom(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitIdAtom"):
                listener.exitIdAtom(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitIdAtom"):
                return visitor.visitIdAtom(self)
            else:
                return visitor.visitChildren(self)


    class StringAtomContext(AtomContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.AtomContext)
            super(AsaliLangGrammarParser.StringAtomContext, self).__init__(parser)
            self.copyFrom(ctx)

        def STRING(self):
            return self.getToken(AsaliLangGrammarParser.STRING, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterStringAtom"):
                listener.enterStringAtom(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitStringAtom"):
                listener.exitStringAtom(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitStringAtom"):
                return visitor.visitStringAtom(self)
            else:
                return visitor.visitChildren(self)


    class NilAtomContext(AtomContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.AtomContext)
            super(AsaliLangGrammarParser.NilAtomContext, self).__init__(parser)
            self.copyFrom(ctx)

        def NIL(self):
            return self.getToken(AsaliLangGrammarParser.NIL, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterNilAtom"):
                listener.enterNilAtom(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitNilAtom"):
                listener.exitNilAtom(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitNilAtom"):
                return visitor.visitNilAtom(self)
            else:
                return visitor.visitChildren(self)


    class NumberAtomContext(AtomContext):

        def __init__(self, parser, ctx): # actually a AsaliLangGrammarParser.AtomContext)
            super(AsaliLangGrammarParser.NumberAtomContext, self).__init__(parser)
            self.copyFrom(ctx)

        def INT(self):
            return self.getToken(AsaliLangGrammarParser.INT, 0)
        def FLOAT(self):
            return self.getToken(AsaliLangGrammarParser.FLOAT, 0)

        def enterRule(self, listener):
            if hasattr(listener, "enterNumberAtom"):
                listener.enterNumberAtom(self)

        def exitRule(self, listener):
            if hasattr(listener, "exitNumberAtom"):
                listener.exitNumberAtom(self)

        def accept(self, visitor):
            if hasattr(visitor, "visitNumberAtom"):
                return visitor.visitNumberAtom(self)
            else:
                return visitor.visitChildren(self)



    def atom(self):

        localctx = AsaliLangGrammarParser.AtomContext(self, self._ctx, self.state)
        self.enterRule(localctx, 34, self.RULE_atom)
        self._la = 0 # Token type
        try:
            self.state = 194
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [20]:
                localctx = AsaliLangGrammarParser.ParExprContext(self, localctx)
                self.enterOuterAlt(localctx, 1)
                self.state = 185
                self.match(AsaliLangGrammarParser.OPAR)
                self.state = 186
                self.expr(0)
                self.state = 187
                self.match(AsaliLangGrammarParser.CPAR)
                pass
            elif token in [38, 39]:
                localctx = AsaliLangGrammarParser.NumberAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 2)
                self.state = 189
                _la = self._input.LA(1)
                if not(_la==38 or _la==39):
                    self._errHandler.recoverInline(self)
                else:
                    self._errHandler.reportMatch(self)
                    self.consume()
                pass
            elif token in [28, 29]:
                localctx = AsaliLangGrammarParser.BooleanAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 3)
                self.state = 190
                _la = self._input.LA(1)
                if not(_la==28 or _la==29):
                    self._errHandler.recoverInline(self)
                else:
                    self._errHandler.reportMatch(self)
                    self.consume()
                pass
            elif token in [37]:
                localctx = AsaliLangGrammarParser.IdAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 4)
                self.state = 191
                self.match(AsaliLangGrammarParser.ID)
                pass
            elif token in [40]:
                localctx = AsaliLangGrammarParser.StringAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 5)
                self.state = 192
                self.match(AsaliLangGrammarParser.STRING)
                pass
            elif token in [30]:
                localctx = AsaliLangGrammarParser.NilAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 6)
                self.state = 193
                self.match(AsaliLangGrammarParser.NIL)
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx



    def sempred(self, localctx, ruleIndex, predIndex):
        if self._predicates == None:
            self._predicates = dict()
        self._predicates[16] = self.expr_sempred
        pred = self._predicates.get(ruleIndex, None)
        if pred is None:
            raise Exception("No predicate with index:" + str(ruleIndex))
        else:
            return pred(localctx, predIndex)

    def expr_sempred(self, localctx, predIndex):
            if predIndex == 0:
                return self.precpred(self._ctx, 10)
         

            if predIndex == 1:
                return self.precpred(self._ctx, 7)
         

            if predIndex == 2:
                return self.precpred(self._ctx, 6)
         

            if predIndex == 3:
                return self.precpred(self._ctx, 5)
         

            if predIndex == 4:
                return self.precpred(self._ctx, 4)
         

            if predIndex == 5:
                return self.precpred(self._ctx, 3)
         

            if predIndex == 6:
                return self.precpred(self._ctx, 2)
         




