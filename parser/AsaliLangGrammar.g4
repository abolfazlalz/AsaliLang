grammar AsaliLangGrammar;

parse
 : block EOF
 ;

block
 : stat*
 ;

stat
 : assignment
 | ifStat
 | whileStat
 | methodCallStat
 | forStat
 | loopStat
 | defineFuncStats
 ;

assignment
 : ID ASSIGN expr SCOL
 ;

ifStat
 : IF conditionBlock (ELSE IF conditionBlock)* (ELSE statBlock)?
 ;

conditionBlock
 : expr statBlock
 ;

statBlock
 : OBRACE block CBRACE
 | DO? BEGIN block END
 | stat
 | THEN block
 ;

whileStat
 : WHILE expr statBlock
 ;

forStat
 : FOR ID ASSIGN expr COL expr statBlock
 ;

loopStat
 : LOOP ID COL expr statBlock
 ;

methodCallStat
 : methodCall SCOL
 ;

defineFuncStats
 : FUNC ID OPAR defineFuncArguments CPAR statBlock
 ;

 methodCall
 : ID methodCallArguments
 ;

 inlineMethodCall
 : ID OPAR methodCallArguments CPAR
 ;

methodCallArguments
 : // No arguments
 | expr (',' expr)*  // Some arguments
 ;

defineFuncArguments
 : // No arguments
 | ID (',' ID)*
 ;


expr
 : inlineMethodCall                       #methodCallExpr
 |<assoc=right> expr POW expr           #powExpr
 | MINUS expr                           #unaryMinusExpr
 | NOT expr                             #notExpr
 | expr op=(MULT | DIV | MOD) expr      #multiplicationExpr
 | expr op=(PLUS | MINUS) expr          #additiveExpr
 | expr op=(LTEQ | GTEQ | LT | GT) expr #relationalExpr
 | expr op=(EQ | NEQ) expr              #equalityExpr
 | expr AND expr                        #andExpr
 | expr OR expr                         #orExpr
 | atom                                 #atomExpr
 ;

atom
 : OPAR expr CPAR #parExpr
 | (INT | FLOAT)  #numberAtom
 | (TRUE | FALSE) #booleanAtom
 | ID             #idAtom
 | STRING         #stringAtom
 | NIL            #nilAtom
 ;

OR : '||';
AND : '&&';
EQ : '==';
NEQ : '!=';
GT : '>';
LT : '<';
GTEQ : '>=';
LTEQ : '<=';
PLUS : '+';
MINUS : '-';
MULT : '*';
DIV : '/';
MOD : '%';
POW : '^';
NOT : '!';

SCOL : ';';
COL : ':';
ASSIGN : '=';
OPAR : '(';
CPAR : ')';
OBRACE : '{';
CBRACE : '}';

BEGIN : 'begin';
END : 'end';
DO : 'do';
THEN : 'then';

TRUE : 'true';
FALSE : 'false';
NIL : 'nil';
IF : 'if';
ELSE : 'else';
WHILE : 'while';
FOR : 'for';
LOOP : 'loop';
FUNC : 'func';

ID
 : [a-zA-Z_] [a-zA-Z_0-9]*
 ;

INT
 : [0-9]+
 ;

FLOAT
 : [0-9]+ '.' [0-9]*
 | '.' [0-9]+
 ;

STRING
 : '"' (~["\r\n] | '""')* '"'
 ;

COMMENT
 : '#' ~[\r\n]* -> skip
 ;

SPACE
 : [ \t\r\n] -> skip
 ;