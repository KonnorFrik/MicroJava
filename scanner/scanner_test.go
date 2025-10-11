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

func TestReadLexem_Name(t *testing.T) {
	reader := litReader.New(strings.NewReader("name"))
	scanner := New(reader)
	wantLex := Lexem {
		Kind: LEX_IDENT,
		Line: 1,
		Col: 1,
		RawVal: "name",
	}
	lex, err := scanner.ReadLexem()

	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_Number(t *testing.T) {
	reader := litReader.New(strings.NewReader("123"))
	scanner := New(reader)
	wantLex := Lexem {
		Kind: LEX_NUMBER,
		Line: 1,
		Col: 1,
		RawVal: "123",
		NumVal: 123,
	}
	lex, err := scanner.ReadLexem()

	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

func TestReadLexem_CharConst(t *testing.T) {
	reader := litReader.New(strings.NewReader("'a'"))
	scanner := New(reader)
	wantLex := Lexem {
		Kind: LEX_CHAR_CON,
		Line: 1,
		Col: 1,
		RawVal: "a",
		NumVal: 97,
	}
	lex, err := scanner.ReadLexem()

	assert.Nil(t, err)
	assert.Equal(t, wantLex, lex)
}

