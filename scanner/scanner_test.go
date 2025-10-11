package scanner

import (
	litReader "mjc/reader"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymbol_Success(t *testing.T) {
	var data litReader.Liter

	for data = LIT_A_SMALL; data <= LIT_Z_SMALL; data++ {
		assert.True(t, isSymbol(data))
	}

	for data = LIT_A_BIG; data <= LIT_Z_BIG; data++ {
		assert.True(t, isSymbol(data))
	}
}

func TestIsSymbol_Number(t *testing.T) {
	var data litReader.Liter

	for data = LIT_0; data <= LIT_9; data++ {
		assert.False(t, isSymbol(data))
	}
}

func TestIsSymbol_EOF(t *testing.T) {
	assert.False(t, isSymbol(litReader.EOF_CHAR))
}

func TestIsSymbol_EOL(t *testing.T) {
	assert.False(t, isSymbol(LIT_EOL))
}

func TestNew_Success(t *testing.T) {
	reader := litReader.New(strings.NewReader("foo"))
	scanner := New(reader)
	assert.NotNil(t, scanner)
}

func TestReadLexem_Idennt(t *testing.T) {
	reader := litReader.New(strings.NewReader("name"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_IDENT,
		Line:   1,
		Col:    1,
		RawVal: "name",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Idennt_WithSpaces(t *testing.T) {
	reader := litReader.New(strings.NewReader("  	name 	 "))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_IDENT,
		Line:   1,
		Col:    4,
		RawVal: "name",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Idennt_TwoLines(t *testing.T) {
	reader := litReader.New(strings.NewReader("foo\nsecondName"))
	scanner := New(reader)
	wantLex1 := Lexem{
		Kind:   LEX_IDENT,
		Line:   1,
		Col:    1,
		RawVal: "foo",
	}
	wantLex2 := Lexem{
		Kind:   LEX_IDENT,
		Line:   2,
		Col:    1,
		RawVal: "secondName",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex1, lex)

	lex, err = scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex2, lex)
}

func TestReadLexem_Ident_KW_break(t *testing.T) {
	reader := litReader.New(strings.NewReader("break"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_BREAK,
		Line:   1,
		Col:    1,
		RawVal: "break",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_class(t *testing.T) {
	reader := litReader.New(strings.NewReader("class"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_CLASS,
		Line:   1,
		Col:    1,
		RawVal: "class",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_else(t *testing.T) {
	reader := litReader.New(strings.NewReader("else"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_ELSE,
		Line:   1,
		Col:    1,
		RawVal: "else",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_final(t *testing.T) {
	reader := litReader.New(strings.NewReader("final"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_FINAL,
		Line:   1,
		Col:    1,
		RawVal: "final",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_if(t *testing.T) {
	reader := litReader.New(strings.NewReader("if"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_IF,
		Line:   1,
		Col:    1,
		RawVal: "if",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_new(t *testing.T) {
	reader := litReader.New(strings.NewReader("new"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_NEW,
		Line:   1,
		Col:    1,
		RawVal: "new",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_print(t *testing.T) {
	reader := litReader.New(strings.NewReader("print"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_PRINT,
		Line:   1,
		Col:    1,
		RawVal: "print",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_program(t *testing.T) {
	reader := litReader.New(strings.NewReader("program"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_PROGRAM,
		Line:   1,
		Col:    1,
		RawVal: "program",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_read(t *testing.T) {
	reader := litReader.New(strings.NewReader("read"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_READ,
		Line:   1,
		Col:    1,
		RawVal: "read",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_return(t *testing.T) {
	reader := litReader.New(strings.NewReader("return"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_RETURN,
		Line:   1,
		Col:    1,
		RawVal: "return",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_void(t *testing.T) {
	reader := litReader.New(strings.NewReader("void"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_VOID,
		Line:   1,
		Col:    1,
		RawVal: "void",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Ident_KW_while(t *testing.T) {
	reader := litReader.New(strings.NewReader("while"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_KW_WHILE,
		Line:   1,
		Col:    1,
		RawVal: "while",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Number(t *testing.T) {
	reader := litReader.New(strings.NewReader("123"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_NUMBER,
		Line:   1,
		Col:    1,
		RawVal: "123",
		NumVal: 123,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Number_TooBig(t *testing.T) {
	reader := litReader.New(strings.NewReader("9999999999999999999999999"))
	scanner := New(reader)
	wantLex := Lexem{
		// Kind: LEX_NUMBER,
		// Line: 1,
		// Col: 1,
		// RawVal: "9999999999999999999999999",
		// NumVal: 0,
	}

	lex, err := scanner.ReadLexem()
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "ScannerError")
	assert.ErrorContains(t, err, "readNumber")
	assert.ErrorContains(t, err, "cannot parse number")
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_CharConst(t *testing.T) {
	reader := litReader.New(strings.NewReader("'a'"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_CHAR_CON,
		Line:   1,
		Col:    1,
		RawVal: "a",
		NumVal: 97,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_CharConst_Invalid(t *testing.T) {
	reader := litReader.New(strings.NewReader("'a"))
	scanner := New(reader)
	wantLex := Lexem{
		// Kind: LEX_CHAR_CON,
		// Line: 1,
		// Col: 1,
		// RawVal: "a",
		// NumVal: 97,
	}

	lex, err := scanner.ReadLexem()
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "ScannerError")
	assert.ErrorContains(t, err, "readCharConst")
	assert.ErrorContains(t, err, "invalid char constant")
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_CharConst_Number(t *testing.T) {
	reader := litReader.New(strings.NewReader("'2'"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_CHAR_CON,
		Line:   1,
		Col:    1,
		RawVal: "2",
		NumVal: 50,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Plus_Once(t *testing.T) {
	reader := litReader.New(strings.NewReader("+"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_PLUS,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Plus_LostChar(t *testing.T) {
	reader := litReader.New(strings.NewReader("+2"))
	scanner := New(reader)
	wantLex1 := Lexem{
		Kind: LEX_PLUS,
		Line: 1,
		Col:  1,
	}
	wantLex2 := Lexem{
		Kind:   LEX_NUMBER,
		Line:   1,
		Col:    2,
		RawVal: "2",
		NumVal: 2,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex1, lex)

	lex, err = scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex2, lex)
}

func TestReadLexem_Plus_Twice(t *testing.T) {
	reader := litReader.New(strings.NewReader("++"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_PPLUS,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Minus_Once(t *testing.T) {
	reader := litReader.New(strings.NewReader("-"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_MINUS,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Minus_Twice(t *testing.T) {
	reader := litReader.New(strings.NewReader("--"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_MMINUS,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Asterisk(t *testing.T) {
	reader := litReader.New(strings.NewReader("*"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_MULT,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Module(t *testing.T) {
	reader := litReader.New(strings.NewReader("%"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_REM,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_NotEq(t *testing.T) {
	reader := litReader.New(strings.NewReader("!="))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_NEQ,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_NotEq_OnlyExclamMark(t *testing.T) {
	reader := litReader.New(strings.NewReader("!"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_NONE,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Less(t *testing.T) {
	reader := litReader.New(strings.NewReader("<"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_LSS,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_LessThen(t *testing.T) {
	reader := litReader.New(strings.NewReader("<="))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_LEQ,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Less_Number(t *testing.T) {
	reader := litReader.New(strings.NewReader("< 1"))
	scanner := New(reader)
	wantLex1 := Lexem{
		Kind: LEX_LSS,
		Line: 1,
		Col:  1,
	}
	wantLex2 := Lexem{
		Kind:   LEX_NUMBER,
		Line:   1,
		Col:    3,
		RawVal: "1",
		NumVal: 1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex1, lex)

	assert.Equal(t, ' ', scanner.char)

	lex, err = scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex2, lex)
}

func TestReadLexem_Greater(t *testing.T) {
	reader := litReader.New(strings.NewReader(">"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_GTR,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_GreaterThen(t *testing.T) {
	reader := litReader.New(strings.NewReader(">="))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_GEQ,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Greater_Number(t *testing.T) {
	reader := litReader.New(strings.NewReader("> 1"))
	scanner := New(reader)
	wantLex1 := Lexem{
		Kind: LEX_GTR,
		Line: 1,
		Col:  1,
	}
	wantLex2 := Lexem{
		Kind:   LEX_NUMBER,
		Line:   1,
		Col:    3,
		RawVal: "1",
		NumVal: 1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex1, lex)

	assert.Equal(t, ' ', scanner.char)

	lex, err = scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex2, lex)
}

func TestReadLexem_Or(t *testing.T) {
	reader := litReader.New(strings.NewReader("||"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_OR,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Or_Invalid(t *testing.T) {
	reader := litReader.New(strings.NewReader("|"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_NONE,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
	assert.Equal(t, rune(litReader.EOF_CHAR), scanner.char)
}

func TestReadLexem_Semicolon(t *testing.T) {
	reader := litReader.New(strings.NewReader(";"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_SEMICOLON,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Comma(t *testing.T) {
	reader := litReader.New(strings.NewReader(","))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_COMMA,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Period(t *testing.T) {
	reader := litReader.New(strings.NewReader("."))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_PERIOD,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_LeftPar(t *testing.T) {
	reader := litReader.New(strings.NewReader("("))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_LPAR,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_RightPar(t *testing.T) {
	reader := litReader.New(strings.NewReader(")"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_RPAR,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_LeftBrack(t *testing.T) {
	reader := litReader.New(strings.NewReader("["))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_LBRACK,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_RightBrack(t *testing.T) {
	reader := litReader.New(strings.NewReader("]"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_RBRACK,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_LeftBrace(t *testing.T) {
	reader := litReader.New(strings.NewReader("{"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_LBRACE,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_RightBrace(t *testing.T) {
	reader := litReader.New(strings.NewReader("}"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_RBRACE,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_EOF(t *testing.T) {
	reader := litReader.New(strings.NewReader(""))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_EOF,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Equal(t *testing.T) {
	reader := litReader.New(strings.NewReader("=="))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_EQL,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Assign(t *testing.T) {
	reader := litReader.New(strings.NewReader("="))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_ASSIGN,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_And(t *testing.T) {
	reader := litReader.New(strings.NewReader("&&"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_AND,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_And_Invalid(t *testing.T) {
	reader := litReader.New(strings.NewReader("&"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_NONE,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Comment_Ident(t *testing.T) {
	reader := litReader.New(strings.NewReader("// this is a comment\nfoo"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind:   LEX_IDENT,
		Line:   2,
		Col:    1,
		RawVal: "foo",
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Slash(t *testing.T) {
	reader := litReader.New(strings.NewReader("/"))
	scanner := New(reader)
	wantLex := Lexem{
		Kind: LEX_SLASH,
		Line: 1,
		Col:  1,
	}

	lex, err := scanner.ReadLexem()
	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}
