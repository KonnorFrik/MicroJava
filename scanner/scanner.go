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
    charNum int
}

func New(input LiterReader) *Scanner {
    var obj Scanner
    obj.reader = input
    obj.readLiter()
    return &obj
}

func (s *Scanner) ReadLexem() Lexem {
}

// read liter from input and store inside of s
func (s *Scanner) readLiter() {
    s.char = s.reader.MustReadLiter()
    s.charNum++

    if s.char == '\n' {
        s.charNum = 0
        s.lineNum++
    }
}
