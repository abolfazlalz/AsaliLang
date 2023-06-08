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
 | OTHER {fmt.Println("unknown char: " + $OTHER.text);}
 ;

assignment
 : ID ASSIGN expr SCOL
 ;

ifStat
 : IF condition_block (ELSE IF condition_block)* (ELSE stat_block)?
 ;

condition_block
 : expr stat_block
 ;

stat_block
 : OBRACE block CBRACE
 | DO? BEGIN block END
 | stat
 | THEN block
 ;

whileStat
 : WHILE expr stat_block
 ;

forStat
 : FOR ID ASSIGN expr COL expr stat_block
 ;

loopStat
 : LOOP ID COL expr stat_block
 ;

methodCallStat
 : methodCall SCOL
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

OTHER
 : .
 ;