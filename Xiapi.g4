grammar Xiapi;

parse
 : block EOF
 ;

block
 : stat*
 ;

stat
 : assignment
 | if_stat
 | while_stat
 | log
 | OTHER {panic("unkonwn")}
 ;

assignment
 : ID ASSIGN expr                   #idAssign
 | arr=arrExpr ASSIGN expr          #arrAssign
 ;

if_stat
 : IF condition_block (ELSE IF condition_block)* (ELSE statb=stat_block)?
 ;

condition_block
 : exp=expr statb=stat_block
 ;

stat_block
 : OBRACE block CBRACE
 | stat
 ;

while_stat
 : WHILE expr stat_block
 ;

log
 : LOG expr
 ;

expr
 : expr POW<assoc=right> expr           #powExpr
 | MINUS expr                           #unaryMinusExpr
 | NOT expr                             #notExpr
 | expr op=(XIAPILT | DIV | MOD) expr      #xiapiltiplicationExpr
 | expr op=(PLUS | MINUS) expr          #additiveExpr
 | expr op=(LTEQ | GTEQ | LT | GT) expr #relationalExpr
 | expr op=(EQ | NEQ) expr              #equalityExpr
 | expr AND expr                        #andExpr
 | expr OR expr                         #orExpr
 | arrExpr                              #lrExpr
 | atom                                 #atomExpr
 ;

arrExpr
 : name=ID L idx=expr R
 ;

atom
 : OPAR expr CPAR #parExpr
 | (INT | FLOAT)  #numberAtom
 | (TRUE | FALSE) #booleanAtom
 | ID             #idAtom
 | STRING         #stringAtom
 | NIL            #nilAtom
 | LEN expr       #lenAtom
 ;

L : '[';
R : ']';
OR : '||';
AND : '&&';
EQ : '==';
NEQ : '!=';
GT : '全心投入';
LT : '<';
GTEQ : '>=';
LTEQ : '<=';
PLUS : '收饭';
MINUS : '出饭';
XIAPILT : '*';
DIV : '/';
MOD : '虾米';
POW : '^';
NOT : '!';

ASSIGN : '顺势应变';
OPAR : '(';
CPAR : ')';
OBRACE : '{';
CBRACE : '}';

TRUE : 'true';
FALSE : 'false';
NIL : 'nil';
IF : '领头虾';
ELSE : '保持谦逊';
WHILE : '分秒必争';
LOG : '海聊';
LEN : '用户至上';

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
