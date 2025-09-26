package reader

import (
	"bufio"
	"io"
)

// Liter - one symbol.
type Liter = rune

// LiterReader - Read liters one by one from given input.
type LiterReader struct {
    reader *bufio.Reader
}

func New(input io.Reader) *LiterReader {
    var obj LiterReader
    obj.reader = bufio.NewReader(input)
    return &obj
}

func (lr *LiterReader) ReadLiter() (Liter, error) {
    char, _, err := lr.reader.ReadRune()
    return char, err
}

// Read one liter or panic at any error.
// Return '-1' as EOF.
func (lr *LiterReader) MustReadLiter() Liter {
    char, _, err := lr.reader.ReadRune()

    if err == io.EOF {
        return -1
    }

    if err != nil {
        panic(err)
    }

    return char
}
