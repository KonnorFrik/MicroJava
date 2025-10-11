package scanner

import "fmt"

type Token struct {
    Kind int      // Lexem type
    Line int      // Where Lexem start
    Col  int      // Where Lexem start
    RawVal string // Raw value from input
    NumVal int    // Converted numbers
}

type Lexem = Token

func (t *Token) NewError(msg string) error {
    return fmt.Errorf(
        "TokenError: %s at line:%d column:%d, for raw value:%s",
        msg,
        t.Line,
        t.Col,
        t.RawVal,
    )
}
