package scanner

import (
	"fmt"
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

	char    litReader.Liter
	lineNum int
	colNum  int
}

const debug = false

func New(input LiterReader) *Scanner {
	var obj Scanner
	obj.reader = input
	obj.lineNum = 1
	obj.char = obj.readLiter()
	return &obj
}

func (s *Scanner) ReadLexem() (Lexem, error) {
	if debug {
		fmt.Printf("ReadLexem: Start with stored char '%c'(%[1]d)\n", s.char)
	}

	for s.char >= 0 && s.char <= ' ' {
		s.char = s.readLiter()

		if debug {
			fmt.Printf("ReadLexem: Skip space symb (%d)\n", s.char)
		}
	}

	var (
		token = Token{
			Line: s.lineNum,
			Col:  s.colNum,
			Kind: LEX_NONE,
		}
		err error
	)

	switch {
	case isSymbol(s.char):
		err = s.readName(&token)

	case isNumber(s.char):
		err = s.readNumber(&token)

	case s.char == LIT_APOSTROPHE:
		err = s.readCharConst(&token)

	case s.char == LIT_PLUS:
		s.char = s.readLiter()

		if s.char == LIT_PLUS {
			token.Kind = LEX_PPLUS

		} else {
			token.Kind = LEX_PLUS
		}

	case s.char == LIT_MINUS:
		s.char = s.readLiter()

		if s.char == LIT_MINUS {
			token.Kind = LEX_MMINUS

		} else {
			token.Kind = LEX_MINUS
		}

	case s.char == LIT_ASTERISK:
		token.Kind = LEX_MULT

	case s.char == LIT_PERCENT:
		token.Kind = LEX_REM

	case s.char == LIT_EXCLAMATION:
		s.char = s.readLiter()

		if s.char == LIT_EQUAL {
			token.Kind = LEX_NEQ
		}

	case s.char == LIT_LESS:
		s.char = s.readLiter()

		if s.char == LIT_EQUAL {
			token.Kind = LEX_LEQ

		} else {
			token.Kind = LEX_LSS
		}

	case s.char == LIT_GREATER:
		s.char = s.readLiter()

		if s.char == LIT_EQUAL {
			token.Kind = LEX_GEQ

		} else {
			token.Kind = LEX_GTR
		}

	case s.char == LIT_VERTICAL_BAR:
		s.char = s.readLiter()

		if s.char == LIT_VERTICAL_BAR {
			token.Kind = LEX_OR
		}

	case s.char == LIT_SEMICOLON:
		token.Kind = LEX_SEMICOLON

	case s.char == LIT_COMMA:
		token.Kind = LEX_COMMA

	case s.char == LIT_PERIOD:
		token.Kind = LEX_PERIOD

	case s.char == LIT_L_PAR:
		token.Kind = LEX_LPAR

	case s.char == LIT_R_PAR:
		token.Kind = LEX_RPAR

	case s.char == LIT_L_BRACK:
		token.Kind = LEX_LBRACK

	case s.char == LIT_R_BRACK:
		token.Kind = LEX_RBRACK

	case s.char == LIT_L_BRACE:
		token.Kind = LEX_LBRACE

	case s.char == LIT_R_BRACE:
		token.Kind = LEX_RBRACE

	case s.char == litReader.EOF_CHAR:
		token.Kind = LEX_EOF

	case s.char == LIT_EQUAL:
		s.char = s.readLiter()

		if s.char == LIT_EQUAL {
			token.Kind = LEX_EQL

		} else {
			token.Kind = LEX_ASSIGN
		}

	case s.char == LIT_AMPERSAND:
		s.char = s.readLiter()

		if s.char == LIT_AMPERSAND {
			token.Kind = LEX_AND
		}

	case s.char == LIT_SLASH:
		s.char = s.readLiter()

		if s.char == LIT_SLASH {
			for s.char != LIT_EOL && s.char != litReader.EOF_CHAR {
				s.char = s.readLiter()
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
		if debug {
			fmt.Printf("readLiter: read EOF at line:%d col:%d\n", s.lineNum, s.colNum)
		}

		return
	}

	if char == LIT_EOL {
		s.colNum = 0
		s.lineNum++
	}

	if debug {
		fmt.Printf("readLiter: read '%c'(%[1]d) at line:%d col:%d\n", char, s.lineNum, s.colNum)
	}
	return
}

var keywords = map[string]LexemKind{
	"break":   LEX_KW_BREAK,
	"class":   LEX_KW_CLASS,
	"else":    LEX_KW_ELSE,
	"final":   LEX_KW_FINAL,
	"if":      LEX_KW_IF,
	"new":     LEX_KW_NEW,
	"print":   LEX_KW_PRINT,
	"program": LEX_KW_PROGRAM,
	"read":    LEX_KW_READ,
	"return":  LEX_KW_RETURN,
	"void":    LEX_KW_VOID,
	"while":   LEX_KW_WHILE,
}

// readName - read a complete sequence of symbols and numbers.
// write Token.Kind = LEX_KW_* or LEX_IDENT.
func (s *Scanner) readName(token *Token) error {
	var name strings.Builder
	name.WriteRune(s.char)
	s.char = s.readLiter()

	for isSymbol(s.char) || isNumber(s.char) {
		name.WriteRune(s.char)
		s.char = s.readLiter()
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
func (s *Scanner) readNumber(token *Token) error {
	token.Kind = LEX_NUMBER
	var number strings.Builder
	number.WriteRune(s.char)
	s.char = s.readLiter()

	for isNumber(s.char) {
		number.WriteRune(s.char)
		s.char = s.readLiter()
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
