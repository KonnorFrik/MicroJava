package scanner

import (
	"fmt"
	litReader "mjc/reader"
	"os"
	"strconv"
	"strings"
)

type Token struct {
    Kind int      // Lexem type
    Line int      // Where Lexem start
    Col  int      // Where Lexem start
    RawVal string // Raw value from input
    NumVal int    // Converted numbers
}

type Lexem = Token

type LiterReader interface {
    ReadLiter() (litReader.Liter, error)
    MustReadLiter() litReader.Liter
}

// Scanner - read input program and parse to Lexems.
// v0.0.1 - Always use 'MustReadLiter'
type Scanner struct {
    reader LiterReader

    char litReader.Liter
    lineNum int
    colNum int
}

func New(input LiterReader) *Scanner {
    var obj Scanner
    obj.reader = input
    obj.readLiter()
    return &obj
}

func (s *Scanner) ReadLexem() Lexem {
    for s.char >= 0 && s.char <= ' ' {
        s.readLiter()
    }

    var token = Token {
        Line: s.lineNum,
        Col: s.colNum,
        Kind: LEX_NONE,
    }

    switch {
    case isSymbol(s.char):
        s.readName(&token)

    case isNumber(s.char):
        s.readNumber(&token)

    case s.char == LIT_APOSTROPHE:
        s.readCharConst(&token)


    case s.char == LIT_PLUS:
        s.readLiter()

        if s.char == LIT_PLUS {
            token.Kind = LEX_PPLUS

        } else {
            token.Kind = LEX_PLUS
        }

    case s.char == LIT_MINUS:
        s.readLiter()

        if s.char == LIT_MINUS {
            token.Kind = LEX_MMINUS

        } else {
            token.Kind = LEX_MINUS
        }

    case s.char == LIT_ASTERISK:
        s.readLiter()
        token.Kind = LEX_TIMES

    case s.char == LIT_PERCENT:
        s.readLiter()
        token.Kind = LEX_REM


    case s.char == LIT_EXCLAMATION:
        s.readLiter()

        if s.char == LIT_EQUAL {
            token.Kind = LEX_NEQ
        }

    case s.char == LIT_LESS:
        s.readLiter()

        if s.char == LIT_EQUAL {
            token.Kind = LEX_LEQ

        } else {
            token.Kind = LEX_LSS
        }

    case s.char == LIT_GREATER:
        s.readLiter()

        if s.char == LIT_EQUAL {
            token.Kind = LEX_GEQ

        } else {
            token.Kind = LEX_GTR
        }

    case s.char == LIT_VERTICAL_BAR:
        s.readLiter()

        if s.char == LIT_EQUAL {
            token.Kind = LEX_OR
        }

    case s.char == LIT_SEMICOLON:
        s.readLiter()
        token.Kind = LEX_SEMICOLON

    case s.char == LIT_COMMA:
        s.readLiter()
        token.Kind = LEX_COMMA

    case s.char == LIT_PERIOD:
        s.readLiter()
        token.Kind = LEX_PERIOD


    case s.char == LIT_L_PAR:
        s.readLiter()
        token.Kind = LEX_LPAR

    case s.char == LIT_R_PAR:
        s.readLiter()
        token.Kind = LEX_RPAR


    case s.char == LIT_L_BRACK:
        s.readLiter()
        token.Kind = LEX_LBRACK

    case s.char == LIT_R_BRACK:
        s.readLiter()
        token.Kind = LEX_RBRACK


    case s.char == LIT_L_BRACE:
        s.readLiter()
        token.Kind = LEX_LBRACE

    case s.char == LIT_R_BRACE:
        s.readLiter()
        token.Kind = LEX_RBRACE


    case s.char == litReader.EOF_CHAR:
        token.Kind = LEX_EOF

    case s.char == LIT_EQUAL:
        s.readLiter()

        if s.char == LIT_EQUAL {
            s.readLiter()
            token.Kind = LEX_EQL

        } else {
            token.Kind = LEX_ASSIGN
        }

    case s.char == LIT_AMPERSAND:
        s.readLiter()

        if s.char == LIT_AMPERSAND {
            s.readLiter()
            token.Kind = LEX_EQL

        }

    case s.char == LIT_SLASH:
        s.readLiter()

        if s.char == LIT_SLASH {
            for s.char != LIT_EOL && s.char != litReader.EOF_CHAR {
                s.readLiter()
            }

            token = s.ReadLexem()

        } else {
            token.Kind = LEX_SLASH
        }
    }

    return token
}

// read liter from input and store inside of s
func (s *Scanner) readLiter() {
    s.char = s.reader.MustReadLiter()
    s.colNum++

    if s.char == litReader.EOF_CHAR {
        return
    }

    if s.char == LIT_EOL {
        s.colNum = 0
        s.lineNum++
    }
}

var keywords = map[string]LexemKind{
    "break": LEX_KW_BREAK,
    "class": LEX_KW_CLASS,
    "else": LEX_KW_ELSE,
    "final": LEX_KW_FINAL,
    "if": LEX_KW_IF,
    "new": LEX_KW_NEW,
    "print": LEX_KW_PRINT,
    "program": LEX_KW_PROGRAM,
    "read": LEX_KW_READ,
    "return": LEX_KW_RETURN,
    "void": LEX_KW_VOID,
    "while": LEX_KW_WHILE,
}

// readName - read a complete sequence of symbols and numbers.
// write Token.Kind = LEX_KW_* or LEX_IDENT.
func (s *Scanner) readName(token *Token) {
    var name strings.Builder
    name.WriteRune(s.char)
    s.readLiter()
    
    for isSymbol(s.char) || isNumber(s.char) {
        name.WriteRune(s.char)
        s.readLiter()
    }

    token.RawVal = name.String()

    if lexem, ok := keywords[token.RawVal]; ok {
        token.Kind = lexem

    } else {
        token.Kind = LEX_IDENT
    }
}

// readNumber - a complete sequence of numbers and convert it into int.
// write Token.Kind = LEX_NUMBER.
// write zero-value (0) if error occured.
func (s *Scanner) readNumber(token *Token) {
    token.Kind = LEX_NUMBER
    var number strings.Builder
    number.WriteRune(s.char)
    s.readLiter()

    for isNumber(s.char) {
        number.WriteRune(s.char)
        s.readLiter()
    }

    // TODO: implement converting by hand (because stdlib return zero-value)
    strNum := number.String()
    token.RawVal = strNum
    num, err := strconv.Atoi(strNum)
    token.NumVal = num

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: readNumber: %s", NewTokenError("cannot parse number", *token))
    }
}

// readCharConst - read all symbols inside the \'...\'.
// write a token.Kind = LEX_CHAR_CON even if error occured.
func (s *Scanner) readCharConst(token *Token) {
}

// isSymbol - Check is liter is in range 'a'-'z' or 'A'-'Z'.
func isSymbol(lit litReader.Liter) bool {
    return (lit >= LIT_A_SMALL && lit <= LIT_Z_SMALL) || (lit >= LIT_A_BIG && lit <= LIT_Z_BIG)
}

func isNumber(lit litReader.Liter) bool {
    return lit >= LIT_0 && lit <= LIT_9
}

func NewTokenError(msg string, token Token) error {
    return fmt.Errorf(
        "TokenError: %s at line:%d column:%d, for raw value:%s",
        msg,
        token.Line,
        token.Col,
        token.RawVal,
    )
}
