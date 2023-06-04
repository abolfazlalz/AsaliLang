grammar MyGrammar;

program : statements;
statements : (statement NEWLINE)*;


statement
    : IDENTIFIER EQ variableSetterTypes #statementDefineVariable
    | 'begin' statements 'end' #statement_begin_end
    | 'if' condition 'then' statement #statement_if
    | 'if' condition 'then' statement 'else' statement #statement_if_else
    | 'while' condition ':' statement #statement_while
    | 'for' IDENTIFIER EQ number COLON number 'do' statement #statement_for
    | 'loop' IDENTIFIER COLON number 'do' statement #statement_loop
    | methodCall #CallMethod
    | 'print' methodCallArguments #StatementPrintMethod
    ;


methodCall
    : IDENTIFIER '(' methodCallArguments ')'
    ;

variableSetterTypes
    : IDENTIFIER
    | methodCall
    | sumExpr
    ;

methodCallArguments
    : // No arguments
    | expression (',' expression)*  // Some arguments
    ;

expression
    : STRING
    | IDENTIFIER
    | methodCall
    | INTEGER
    ;

STRING
    : '"' ~('"')* '"'
    ;

condition :
    sumExpr CONDITION_OP sumExpr #condition_def
    ;

powerExpr : number #powerExprDefault
    | number POWERBY powerExpr #powerExprPower
    ;

multipleExpr : powerExpr #multipleExprDefault
    | multipleExpr MULTI powerExpr #multipleExprMulti
    | multipleExpr DIVIDE powerExpr #multipleExprDivide
    ;

sumExpr : multipleExpr #sumExprDefault
    | sumExpr PLUS multipleExpr #sumExprPlus
    | sumExpr MINUS multipleExpr #sumExprMinus
    ;

number : INTEGER #numberDefault
    | MINUS number #NumberMinus
    | IDENTIFIER #NumberIdentifier
    | OPEN_PARAN sumExpr CLOSE_PARAN #NumberParentheses
    ;

IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]*;

INTEGER : [0-9]+;
EQ : '=';

COLON : ':';
COMMA : ',';

PLUS : '+';
MINUS : '-';
MULTI : '*';
DIVIDE : '/';
POWERBY : '^';

EQUALBY : '==';
NOTEQUALBY : '!=';
LT : '<';
GT : '>';

COT : '"';

OPEN_PARAN : '(';
CLOSE_PARAN : ')';

CONDITION_OP : LT | GT | EQUALBY | NOTEQUALBY;

NEWLINE : [;][\r\n]+;

NEXT_PARAM : [,]+;

EMPTY : [ \t] -> channel(HIDDEN);