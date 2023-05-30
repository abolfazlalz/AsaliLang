grammar MyGrammar;

program : statements;
statements : (statement NEWLINE)*;


statement : IDENTIFIER EQ priority2 #statement_def_var
    | 'begin' statements 'end' #statement_begin_end
    | 'if' condition 'then' statement #statement_if
    | 'if' condition 'then' statement 'else' statement #statement_if_else
    | 'while' condition 'do' statement #statement_while
    | 'for' IDENTIFIER EQ number COLON number 'do' statement #statement_for
    | 'loop' IDENTIFIER COLON number 'do' statement #statement_loop
    | 'print' IDENTIFIER #statement_print
    | 'print' STRING_LITERAL COMMA IDENTIFIER #statement_print_1
    ;

//expr : expr binop expr
////    | ! expr
//    | ( expr)
//    | identifier
//    | NUMBER
//    ;

condition :
    priority1 CONDITION_OP priority1 #condition_def
    ;

priority1 : number #priority1_number
    | priority1 MULTI number #priority1_multiple
    | priority1 DIVIDE number #priority1_divide
    ;

priority2 : priority1 #priority2_def
    | priority2 PLUS priority1 #priority2_plus
    | priority2 MINUS priority1 #priority2_minus
    ;

number : INTEGER #number_def
    | MINUS number #number_minus
    | IDENTIFIER #number_identifier
    | OPEN_PARAN priority1 CLOSE_PARAN #number_paran
    ;

IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]*;

STRING_LITERAL : ["][*]*["];

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

OPEN_PARAN : '(';
CLOSE_PARAN : ')';

CONDITION_OP : LT | GT | EQUALBY | NOTEQUALBY;

NEWLINE : [\r\n]+;

EMPTY : [ \t] -> channel(HIDDEN);