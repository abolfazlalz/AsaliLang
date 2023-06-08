# Generated from AsaliLangGrammar.g4 by ANTLR 4.13.0
# encoding: utf-8
from __future__ import print_function
from antlr4 import *
from io import StringIO
import sys

def serializedATN():
    return [
        4,1,41,174,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,
        6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,
        2,14,7,14,2,15,7,15,1,0,1,0,1,0,1,1,5,1,37,8,1,10,1,12,1,40,9,1,
        1,2,1,2,1,2,1,2,1,2,1,2,3,2,48,8,2,1,3,1,3,1,3,1,3,1,3,1,4,1,4,1,
        4,1,4,1,4,5,4,60,8,4,10,4,12,4,63,9,4,1,4,1,4,3,4,67,8,4,1,5,1,5,
        1,5,1,6,1,6,1,6,1,6,1,6,3,6,77,8,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,3,
        6,86,8,6,1,7,1,7,1,7,1,7,1,8,1,8,1,8,1,8,1,8,1,8,1,8,1,8,1,9,1,9,
        1,9,1,9,1,9,1,9,1,10,1,10,1,10,1,11,1,11,1,11,1,12,1,12,1,12,1,12,
        1,12,1,13,1,13,1,13,1,13,5,13,121,8,13,10,13,12,13,124,9,13,3,13,
        126,8,13,1,14,1,14,1,14,1,14,1,14,1,14,1,14,3,14,135,8,14,1,14,1,
        14,1,14,1,14,1,14,1,14,1,14,1,14,1,14,1,14,1,14,1,14,1,14,1,14,1,
        14,1,14,1,14,1,14,1,14,1,14,1,14,5,14,158,8,14,10,14,12,14,161,9,
        14,1,15,1,15,1,15,1,15,1,15,1,15,1,15,1,15,1,15,3,15,172,8,15,1,
        15,0,1,28,16,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,0,6,1,0,
        12,14,1,0,10,11,1,0,6,9,1,0,4,5,1,0,37,38,1,0,28,29,186,0,32,1,0,
        0,0,2,38,1,0,0,0,4,47,1,0,0,0,6,49,1,0,0,0,8,54,1,0,0,0,10,68,1,
        0,0,0,12,85,1,0,0,0,14,87,1,0,0,0,16,91,1,0,0,0,18,99,1,0,0,0,20,
        105,1,0,0,0,22,108,1,0,0,0,24,111,1,0,0,0,26,125,1,0,0,0,28,134,
        1,0,0,0,30,171,1,0,0,0,32,33,3,2,1,0,33,34,5,0,0,1,34,1,1,0,0,0,
        35,37,3,4,2,0,36,35,1,0,0,0,37,40,1,0,0,0,38,36,1,0,0,0,38,39,1,
        0,0,0,39,3,1,0,0,0,40,38,1,0,0,0,41,48,3,6,3,0,42,48,3,8,4,0,43,
        48,3,14,7,0,44,48,3,20,10,0,45,48,3,16,8,0,46,48,3,18,9,0,47,41,
        1,0,0,0,47,42,1,0,0,0,47,43,1,0,0,0,47,44,1,0,0,0,47,45,1,0,0,0,
        47,46,1,0,0,0,48,5,1,0,0,0,49,50,5,36,0,0,50,51,5,19,0,0,51,52,3,
        28,14,0,52,53,5,17,0,0,53,7,1,0,0,0,54,55,5,31,0,0,55,61,3,10,5,
        0,56,57,5,32,0,0,57,58,5,31,0,0,58,60,3,10,5,0,59,56,1,0,0,0,60,
        63,1,0,0,0,61,59,1,0,0,0,61,62,1,0,0,0,62,66,1,0,0,0,63,61,1,0,0,
        0,64,65,5,32,0,0,65,67,3,12,6,0,66,64,1,0,0,0,66,67,1,0,0,0,67,9,
        1,0,0,0,68,69,3,28,14,0,69,70,3,12,6,0,70,11,1,0,0,0,71,72,5,22,
        0,0,72,73,3,2,1,0,73,74,5,23,0,0,74,86,1,0,0,0,75,77,5,26,0,0,76,
        75,1,0,0,0,76,77,1,0,0,0,77,78,1,0,0,0,78,79,5,24,0,0,79,80,3,2,
        1,0,80,81,5,25,0,0,81,86,1,0,0,0,82,86,3,4,2,0,83,84,5,27,0,0,84,
        86,3,2,1,0,85,71,1,0,0,0,85,76,1,0,0,0,85,82,1,0,0,0,85,83,1,0,0,
        0,86,13,1,0,0,0,87,88,5,33,0,0,88,89,3,28,14,0,89,90,3,12,6,0,90,
        15,1,0,0,0,91,92,5,34,0,0,92,93,5,36,0,0,93,94,5,19,0,0,94,95,3,
        28,14,0,95,96,5,18,0,0,96,97,3,28,14,0,97,98,3,12,6,0,98,17,1,0,
        0,0,99,100,5,35,0,0,100,101,5,36,0,0,101,102,5,18,0,0,102,103,3,
        28,14,0,103,104,3,12,6,0,104,19,1,0,0,0,105,106,3,22,11,0,106,107,
        5,17,0,0,107,21,1,0,0,0,108,109,5,36,0,0,109,110,3,26,13,0,110,23,
        1,0,0,0,111,112,5,36,0,0,112,113,5,20,0,0,113,114,3,26,13,0,114,
        115,5,21,0,0,115,25,1,0,0,0,116,126,1,0,0,0,117,122,3,28,14,0,118,
        119,5,1,0,0,119,121,3,28,14,0,120,118,1,0,0,0,121,124,1,0,0,0,122,
        120,1,0,0,0,122,123,1,0,0,0,123,126,1,0,0,0,124,122,1,0,0,0,125,
        116,1,0,0,0,125,117,1,0,0,0,126,27,1,0,0,0,127,128,6,14,-1,0,128,
        135,3,24,12,0,129,130,5,11,0,0,130,135,3,28,14,9,131,132,5,16,0,
        0,132,135,3,28,14,8,133,135,3,30,15,0,134,127,1,0,0,0,134,129,1,
        0,0,0,134,131,1,0,0,0,134,133,1,0,0,0,135,159,1,0,0,0,136,137,10,
        10,0,0,137,138,5,15,0,0,138,158,3,28,14,10,139,140,10,7,0,0,140,
        141,7,0,0,0,141,158,3,28,14,8,142,143,10,6,0,0,143,144,7,1,0,0,144,
        158,3,28,14,7,145,146,10,5,0,0,146,147,7,2,0,0,147,158,3,28,14,6,
        148,149,10,4,0,0,149,150,7,3,0,0,150,158,3,28,14,5,151,152,10,3,
        0,0,152,153,5,3,0,0,153,158,3,28,14,4,154,155,10,2,0,0,155,156,5,
        2,0,0,156,158,3,28,14,3,157,136,1,0,0,0,157,139,1,0,0,0,157,142,
        1,0,0,0,157,145,1,0,0,0,157,148,1,0,0,0,157,151,1,0,0,0,157,154,
        1,0,0,0,158,161,1,0,0,0,159,157,1,0,0,0,159,160,1,0,0,0,160,29,1,
        0,0,0,161,159,1,0,0,0,162,163,5,20,0,0,163,164,3,28,14,0,164,165,
        5,21,0,0,165,172,1,0,0,0,166,172,7,4,0,0,167,172,7,5,0,0,168,172,
        5,36,0,0,169,172,5,39,0,0,170,172,5,30,0,0,171,162,1,0,0,0,171,166,
        1,0,0,0,171,167,1,0,0,0,171,168,1,0,0,0,171,169,1,0,0,0,171,170,
        1,0,0,0,172,31,1,0,0,0,12,38,47,61,66,76,85,122,125,134,157,159,
        171
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
                     u"'if'", u"'else'", u"'while'", u"'for'", u"'loop'" ]

    symbolicNames = [ u"<INVALID>", u"<INVALID>", u"OR", u"AND", u"EQ", 
                      u"NEQ", u"GT", u"LT", u"GTEQ", u"LTEQ", u"PLUS", u"MINUS", 
                      u"MULT", u"DIV", u"MOD", u"POW", u"NOT", u"SCOL", 
                      u"COL", u"ASSIGN", u"OPAR", u"CPAR", u"OBRACE", u"CBRACE", 
                      u"BEGIN", u"END", u"DO", u"THEN", u"TRUE", u"FALSE", 
                      u"NIL", u"IF", u"ELSE", u"WHILE", u"FOR", u"LOOP", 
                      u"ID", u"INT", u"FLOAT", u"STRING", u"COMMENT", u"SPACE" ]

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
    RULE_methodCall = 11
    RULE_inlineMethodCall = 12
    RULE_methodCallArguments = 13
    RULE_expr = 14
    RULE_atom = 15

    ruleNames =  [ u"parse", u"block", u"stat", u"assignment", u"ifStat", 
                   u"conditionBlock", u"statBlock", u"whileStat", u"forStat", 
                   u"loopStat", u"methodCallStat", u"methodCall", u"inlineMethodCall", 
                   u"methodCallArguments", u"expr", u"atom" ]

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
    ID=36
    INT=37
    FLOAT=38
    STRING=39
    COMMENT=40
    SPACE=41

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
            self.state = 32
            self.block()
            self.state = 33
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
            self.state = 38
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,0,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    self.state = 35
                    self.stat() 
                self.state = 40
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
            self.state = 47
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,1,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 41
                self.assignment()
                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 42
                self.ifStat()
                pass

            elif la_ == 3:
                self.enterOuterAlt(localctx, 3)
                self.state = 43
                self.whileStat()
                pass

            elif la_ == 4:
                self.enterOuterAlt(localctx, 4)
                self.state = 44
                self.methodCallStat()
                pass

            elif la_ == 5:
                self.enterOuterAlt(localctx, 5)
                self.state = 45
                self.forStat()
                pass

            elif la_ == 6:
                self.enterOuterAlt(localctx, 6)
                self.state = 46
                self.loopStat()
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
            self.state = 49
            self.match(AsaliLangGrammarParser.ID)
            self.state = 50
            self.match(AsaliLangGrammarParser.ASSIGN)
            self.state = 51
            self.expr(0)
            self.state = 52
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
            self.state = 54
            self.match(AsaliLangGrammarParser.IF)
            self.state = 55
            self.conditionBlock()
            self.state = 61
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,2,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    self.state = 56
                    self.match(AsaliLangGrammarParser.ELSE)
                    self.state = 57
                    self.match(AsaliLangGrammarParser.IF)
                    self.state = 58
                    self.conditionBlock() 
                self.state = 63
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,2,self._ctx)

            self.state = 66
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,3,self._ctx)
            if la_ == 1:
                self.state = 64
                self.match(AsaliLangGrammarParser.ELSE)
                self.state = 65
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
            self.state = 68
            self.expr(0)
            self.state = 69
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
            self.state = 85
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [22]:
                self.enterOuterAlt(localctx, 1)
                self.state = 71
                self.match(AsaliLangGrammarParser.OBRACE)
                self.state = 72
                self.block()
                self.state = 73
                self.match(AsaliLangGrammarParser.CBRACE)
                pass
            elif token in [24, 26]:
                self.enterOuterAlt(localctx, 2)
                self.state = 76
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if _la==26:
                    self.state = 75
                    self.match(AsaliLangGrammarParser.DO)


                self.state = 78
                self.match(AsaliLangGrammarParser.BEGIN)
                self.state = 79
                self.block()
                self.state = 80
                self.match(AsaliLangGrammarParser.END)
                pass
            elif token in [31, 33, 34, 35, 36]:
                self.enterOuterAlt(localctx, 3)
                self.state = 82
                self.stat()
                pass
            elif token in [27]:
                self.enterOuterAlt(localctx, 4)
                self.state = 83
                self.match(AsaliLangGrammarParser.THEN)
                self.state = 84
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
            self.state = 87
            self.match(AsaliLangGrammarParser.WHILE)
            self.state = 88
            self.expr(0)
            self.state = 89
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
            self.state = 91
            self.match(AsaliLangGrammarParser.FOR)
            self.state = 92
            self.match(AsaliLangGrammarParser.ID)
            self.state = 93
            self.match(AsaliLangGrammarParser.ASSIGN)
            self.state = 94
            self.expr(0)
            self.state = 95
            self.match(AsaliLangGrammarParser.COL)
            self.state = 96
            self.expr(0)
            self.state = 97
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
            self.state = 99
            self.match(AsaliLangGrammarParser.LOOP)
            self.state = 100
            self.match(AsaliLangGrammarParser.ID)
            self.state = 101
            self.match(AsaliLangGrammarParser.COL)
            self.state = 102
            self.expr(0)
            self.state = 103
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
            self.state = 105
            self.methodCall()
            self.state = 106
            self.match(AsaliLangGrammarParser.SCOL)
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
        self.enterRule(localctx, 22, self.RULE_methodCall)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 108
            self.match(AsaliLangGrammarParser.ID)
            self.state = 109
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
        self.enterRule(localctx, 24, self.RULE_inlineMethodCall)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 111
            self.match(AsaliLangGrammarParser.ID)
            self.state = 112
            self.match(AsaliLangGrammarParser.OPAR)
            self.state = 113
            self.methodCallArguments()
            self.state = 114
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
        self.enterRule(localctx, 26, self.RULE_methodCallArguments)
        self._la = 0 # Token type
        try:
            self.state = 125
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [17, 21]:
                self.enterOuterAlt(localctx, 1)

                pass
            elif token in [11, 16, 20, 28, 29, 30, 36, 37, 38, 39]:
                self.enterOuterAlt(localctx, 2)
                self.state = 117
                self.expr(0)
                self.state = 122
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while _la==1:
                    self.state = 118
                    self.match(AsaliLangGrammarParser.T__0)
                    self.state = 119
                    self.expr(0)
                    self.state = 124
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
        _startState = 28
        self.enterRecursionRule(localctx, 28, self.RULE_expr, _p)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 134
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,8,self._ctx)
            if la_ == 1:
                localctx = AsaliLangGrammarParser.MethodCallExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx

                self.state = 128
                self.inlineMethodCall()
                pass

            elif la_ == 2:
                localctx = AsaliLangGrammarParser.UnaryMinusExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 129
                self.match(AsaliLangGrammarParser.MINUS)
                self.state = 130
                self.expr(9)
                pass

            elif la_ == 3:
                localctx = AsaliLangGrammarParser.NotExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 131
                self.match(AsaliLangGrammarParser.NOT)
                self.state = 132
                self.expr(8)
                pass

            elif la_ == 4:
                localctx = AsaliLangGrammarParser.AtomExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 133
                self.atom()
                pass


            self._ctx.stop = self._input.LT(-1)
            self.state = 159
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,10,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    self.state = 157
                    self._errHandler.sync(self)
                    la_ = self._interp.adaptivePredict(self._input,9,self._ctx)
                    if la_ == 1:
                        localctx = AsaliLangGrammarParser.PowExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 136
                        if not self.precpred(self._ctx, 10):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 10)")
                        self.state = 137
                        self.match(AsaliLangGrammarParser.POW)
                        self.state = 138
                        self.expr(10)
                        pass

                    elif la_ == 2:
                        localctx = AsaliLangGrammarParser.MultiplicationExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 139
                        if not self.precpred(self._ctx, 7):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 7)")
                        self.state = 140
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 28672) != 0)):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 141
                        self.expr(8)
                        pass

                    elif la_ == 3:
                        localctx = AsaliLangGrammarParser.AdditiveExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 142
                        if not self.precpred(self._ctx, 6):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 6)")
                        self.state = 143
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not(_la==10 or _la==11):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 144
                        self.expr(7)
                        pass

                    elif la_ == 4:
                        localctx = AsaliLangGrammarParser.RelationalExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 145
                        if not self.precpred(self._ctx, 5):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 5)")
                        self.state = 146
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 960) != 0)):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 147
                        self.expr(6)
                        pass

                    elif la_ == 5:
                        localctx = AsaliLangGrammarParser.EqualityExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 148
                        if not self.precpred(self._ctx, 4):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 4)")
                        self.state = 149
                        localctx.op = self._input.LT(1)
                        _la = self._input.LA(1)
                        if not(_la==4 or _la==5):
                            localctx.op = self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 150
                        self.expr(5)
                        pass

                    elif la_ == 6:
                        localctx = AsaliLangGrammarParser.AndExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 151
                        if not self.precpred(self._ctx, 3):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 3)")
                        self.state = 152
                        self.match(AsaliLangGrammarParser.AND)
                        self.state = 153
                        self.expr(4)
                        pass

                    elif la_ == 7:
                        localctx = AsaliLangGrammarParser.OrExprContext(self, AsaliLangGrammarParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 154
                        if not self.precpred(self._ctx, 2):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 2)")
                        self.state = 155
                        self.match(AsaliLangGrammarParser.OR)
                        self.state = 156
                        self.expr(3)
                        pass

             
                self.state = 161
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,10,self._ctx)

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
        self.enterRule(localctx, 30, self.RULE_atom)
        self._la = 0 # Token type
        try:
            self.state = 171
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [20]:
                localctx = AsaliLangGrammarParser.ParExprContext(self, localctx)
                self.enterOuterAlt(localctx, 1)
                self.state = 162
                self.match(AsaliLangGrammarParser.OPAR)
                self.state = 163
                self.expr(0)
                self.state = 164
                self.match(AsaliLangGrammarParser.CPAR)
                pass
            elif token in [37, 38]:
                localctx = AsaliLangGrammarParser.NumberAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 2)
                self.state = 166
                _la = self._input.LA(1)
                if not(_la==37 or _la==38):
                    self._errHandler.recoverInline(self)
                else:
                    self._errHandler.reportMatch(self)
                    self.consume()
                pass
            elif token in [28, 29]:
                localctx = AsaliLangGrammarParser.BooleanAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 3)
                self.state = 167
                _la = self._input.LA(1)
                if not(_la==28 or _la==29):
                    self._errHandler.recoverInline(self)
                else:
                    self._errHandler.reportMatch(self)
                    self.consume()
                pass
            elif token in [36]:
                localctx = AsaliLangGrammarParser.IdAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 4)
                self.state = 168
                self.match(AsaliLangGrammarParser.ID)
                pass
            elif token in [39]:
                localctx = AsaliLangGrammarParser.StringAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 5)
                self.state = 169
                self.match(AsaliLangGrammarParser.STRING)
                pass
            elif token in [30]:
                localctx = AsaliLangGrammarParser.NilAtomContext(self, localctx)
                self.enterOuterAlt(localctx, 6)
                self.state = 170
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
        self._predicates[14] = self.expr_sempred
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
         




