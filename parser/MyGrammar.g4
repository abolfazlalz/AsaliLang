grammar MyGrammar;

program : statements EOF;
statements : statement*;


statement
    : assignmentStatement #assignmentStatementBlock
    | ifStatement #ifStatementBlock
    | whileStatement #whileStatementBlock
    | methodCall #methodCallBlock
    | printStatement #printStatementBlock
    ;

assignmentStatement
    : IDENTIFIER ASSIGN variableSetterTypes SEMICOLON
    ;

printStatement
    : PRINT variableSetterTypes
    ;

blockStatement
    : BEGIN statements END
    ;

whileStatement
    : WHILE expression DO statement END
    ;

methodCall
    : IDENTIFIER OPARAN methodCallArguments CPARAN SEMICOLON
    ;

ifStatement
    : IF expression THEN statement;

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
    | sumExpr
    ;

STRING
    : '"' ~('"')* '"'
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
    | OPARAN sumExpr CPARAN #NumberParentheses
    ;

IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]*;

INTEGER : [0-9]+;

ASSIGN : '=';

COLON : ':';
COMMA : ',';

PLUS : '+';
MINUS : '-';
MULTI : '*';
DIVIDE : '/';
POWERBY : '^';

IF : 'if';
ElSE : 'else';
DO : 'do';
BEGIN : 'begin';
END : 'end';
THEN : 'then';
WHILE : 'while';
FOR : 'for';
LOOP : 'loop';
PRINT : 'print';

TRUE : 'true';
FALSE : 'false';
EQ : '==';
NOTEQUALBY : '!=';
LT : '<';
GT : '>';

COT : '"';

SEMICOLON : ';';
OPARAN : '(';
CPARAN : ')';
OBRAC : '{';
CBRAC : '}';


CONDITION_OP : LT | GT | EQ | NOTEQUALBY;

NEXT_PARAM : [,]+;

EMPTY : [ \t\n] -> channel(HIDDEN);