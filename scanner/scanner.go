package scanner

import (
    litReader "mjc/reader"
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
    case (s.char >= LIT_A_SMALL && s.char <= LIT_Z_SMALL) || (s.char >= LIT_A_BIG && s.char <= LIT_Z_BIG):
        s.readName(&token)

    case s.char >= LIT_0 && s.char <= LIT_9:
        s.readNumber(&token)

    case s.char == LIT_APOSTROPHE:
        s.readCharConst(&token)

    case s.char == LIT_SEMICOLON:
        s.readLiter()
        token.Kind = LEX_SEMICOLON

    case s.char == LIT_PERIOD:
        s.readLiter()
        token.Kind = LEX_PERIOD

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

    // TODO: implement for other symbols (or, ...)

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

func (s *Scanner) readNumber(token *Token) {
}

func (s *Scanner) readName(token *Token) {
}

func (s *Scanner) readCharConst(token *Token) {
}
