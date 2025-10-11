package scanner

import (
	"errors"
	"fmt"
)

var (
	// ErrScanner - any error with scanner package
	ErrScanner = errors.New("scanner")
)

func NewScannerError(funcName string) error {
	return fmt.Errorf("ScannerError: %s", funcName)
}

func WrapError(base error, toWrap ...error) error {
	return fmt.Errorf("%w:\n%w", base, errors.Join(toWrap...))
}
