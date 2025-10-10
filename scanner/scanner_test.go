package scanner

import (
	"testing"
	"github.com/stretchr/testify/assert"
	litReader "mjc/reader"
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
