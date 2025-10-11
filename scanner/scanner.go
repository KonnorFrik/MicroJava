package scanner

import (
	"strconv"
	"strings"

	litReader "mjc/reader"
)

type LiterReader interface {
    ReadLiter() (litReader.Liter, error)
    MustReadLiter() litReader.Liter
}

// Scanner - read input program and parse to Lexems.
// v0.0.1 - Always use 'MustReadLiter'
type Scanner struct {
    reader LiterReader

    lineNum int
    colNum int
}


func New(input LiterReader) *Scanner {
    var obj Scanner
    obj.reader = input
	obj.lineNum = 1
    return &obj
}

func (s *Scanner) ReadLexem() (Lexem, error) {
	var char = s.readLiter()

    for char >= 0 && char <= ' ' {
        char = s.readLiter()
    }

    var (
		token = Token {
			Line: s.lineNum,
			Col: s.colNum,
			Kind: LEX_NONE,
		}
		err error
	)

    switch {
    case isSymbol(char):
        err = s.readName(char, &token)

    case isNumber(char):
        err = s.readNumber(char, &token)

    case char == LIT_APOSTROPHE:
        err = s.readCharConst(&token)


    case char == LIT_PLUS:
        char = s.readLiter()

        if char == LIT_PLUS {
            token.Kind = LEX_PPLUS

        } else {
            token.Kind = LEX_PLUS
        }

    case char == LIT_MINUS:
        char = s.readLiter()

        if char == LIT_MINUS {
            token.Kind = LEX_MMINUS

        } else {
            token.Kind = LEX_MINUS
        }

    case char == LIT_ASTERISK:
        token.Kind = LEX_TIMES

    case char == LIT_PERCENT:
        token.Kind = LEX_REM


    case char == LIT_EXCLAMATION:
        char = s.readLiter()

        if char == LIT_EQUAL {
            token.Kind = LEX_NEQ
        }

    case char == LIT_LESS:
        char = s.readLiter()

        if char == LIT_EQUAL {
            token.Kind = LEX_LEQ

        } else {
            token.Kind = LEX_LSS
        }

    case char == LIT_GREATER:
        char = s.readLiter()

        if char == LIT_EQUAL {
            token.Kind = LEX_GEQ

        } else {
            token.Kind = LEX_GTR
        }

    case char == LIT_VERTICAL_BAR:
        char = s.readLiter()

        if char == LIT_EQUAL {
            token.Kind = LEX_OR
        }

    case char == LIT_SEMICOLON:
        token.Kind = LEX_SEMICOLON

    case char == LIT_COMMA:
        token.Kind = LEX_COMMA

    case char == LIT_PERIOD:
        token.Kind = LEX_PERIOD


    case char == LIT_L_PAR:
        token.Kind = LEX_LPAR

    case char == LIT_R_PAR:
        token.Kind = LEX_RPAR


    case char == LIT_L_BRACK:
        token.Kind = LEX_LBRACK

    case char == LIT_R_BRACK:
        token.Kind = LEX_RBRACK


    case char == LIT_L_BRACE:
        token.Kind = LEX_LBRACE

    case char == LIT_R_BRACE:
        token.Kind = LEX_RBRACE


    case char == litReader.EOF_CHAR:
        token.Kind = LEX_EOF

    case char == LIT_EQUAL:
		char = s.readLiter()

        if char == LIT_EQUAL {
            token.Kind = LEX_EQL

        } else {
            token.Kind = LEX_ASSIGN
        }

    case char == LIT_AMPERSAND:
        char = s.readLiter()

        if char == LIT_AMPERSAND {
            token.Kind = LEX_EQL
        }

    case char == LIT_SLASH:
        char = s.readLiter()

        if char == LIT_SLASH {
            for char != LIT_EOL && char != litReader.EOF_CHAR {
                char = s.readLiter()
            }

            token, err = s.ReadLexem()

        } else {
            token.Kind = LEX_SLASH
        }
    }

	if err != nil {
		return Lexem{}, err
	}

    return token, nil
}

// read liter from input and store in 's.char'
func (s *Scanner) readLiter() (char litReader.Liter) {
    char = s.reader.MustReadLiter()
    s.colNum++

    if char == litReader.EOF_CHAR {
        return
    }

    if char == LIT_EOL {
        s.colNum = 1
        s.lineNum++
    }

	return
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
func (s *Scanner) readName(char litReader.Liter, token *Token) error {
    var name strings.Builder
    name.WriteRune(char)
    char = s.readLiter()
    
    for isSymbol(char) || isNumber(char) {
        name.WriteRune(char)
        char = s.readLiter()
    }

    token.RawVal = name.String()

    if lexem, ok := keywords[token.RawVal]; ok {
        token.Kind = lexem

    } else {
        token.Kind = LEX_IDENT
    }

	return nil
}

// readNumber - a complete sequence of numbers and convert it into int.
// write Token.Kind = LEX_NUMBER.
// write zero-value (0) if error occured.
func (s *Scanner) readNumber(char litReader.Liter, token *Token) error {
    token.Kind = LEX_NUMBER
    var number strings.Builder
    number.WriteRune(char)
    char = s.readLiter()

    for isNumber(char) {
        number.WriteRune(char)
        char = s.readLiter()
    }

    // TODO: implement converting by hand (because stdlib return zero-value)
    strNum := number.String()
    token.RawVal = strNum
    num, err := strconv.Atoi(strNum)
    token.NumVal = num

    if err != nil {
		return WrapError(NewScannerError("readNumber"), token.NewError("cannot parse number"), err)
    }

	return nil
}

// readCharConst - read all symbols inside the \'...\'.
// write a token.Kind = LEX_CHAR_CON even if error occured.
func (s *Scanner) readCharConst(token *Token) error {
	// TODO: implement more complex parsing for 
	// '\u4242'
	token.Kind = LEX_CHAR_CON
	var rawName strings.Builder

	var char = s.readLiter()
	token.NumVal = int(char)
	rawName.WriteRune(char)
	char = s.readLiter()

	if char != LIT_APOSTROPHE {
		return WrapError(NewScannerError("readCharConst"), token.NewError("invalid char constant"))
	}

	token.RawVal = rawName.String()

	return nil
}

// isSymbol - Check is liter is in range 'a'-'z' or 'A'-'Z'.
func isSymbol(lit litReader.Liter) bool {
    return (lit >= LIT_A_SMALL && lit <= LIT_Z_SMALL) || (lit >= LIT_A_BIG && lit <= LIT_Z_BIG)
}

func isNumber(lit litReader.Liter) bool {
    return lit >= LIT_0 && lit <= LIT_9
}
