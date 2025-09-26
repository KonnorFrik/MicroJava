package scanner

const (
    // error
    LEX_NONE = iota

    // lexem classes
    LEX_IDENT
    LEX_NUMBER
    LEX_CHAR_CON

    // operators and spec symbols
    LEX_PLUS
    LEX_MINUS
    // multiply
    LEX_TIMES 
    LEX_SLASH
    // module divide
    LEX_REM
    LEX_EQL
    LEX_NEQ
    // LT - less than
    LEX_LSS
    LEX_LEQ
    LEX_GTR
    LEX_GEQ
    LEX_AND
    LEX_OR
    LEX_ASSIGN
    // increment
    LEX_PPLUS
    // decrement
    LEX_MMINUS
    LEX_SEMICOLON
    LEX_COMMA
    // dot - '.'
    LEX_PERIOD
    // (
    LEX_LPAR
    // )
    LEX_RPAR
    // [
    LEX_LBRACK
    // ]
    LEX_RBRACK
    // {
    LEX_LBRACE
    // }
    LEX_RBRACE

    // key words
    LEX_KW_BREAK
    LEX_KW_CLASS
    LEX_KW_ELSE
    LEX_KW_FINAL
    LEX_KW_IF
    LEX_KW_NEW
    LEX_KW_PRINT
    LEX_KW_PROGRAM
    LEX_KW_READ
    LEX_KW_RETURN
    LEX_KW_VOID
    LEX_KW_WHILE

    LEX_EOF
)
