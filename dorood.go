package rook

import (
	"github.com/reiver/go-rook/internal/token/dclr"

	"fmt"
	"io"
)

// Dorood writes to ‘w’ the DOROOD DECLARATION.
//
// I.e.,:
//
//	"!DOROOD\u0085"
func Dorood(w io.Writer) error {
	if nil == w {
		return nil
	}

	{
		str := dclr.Dorood().String()

		n, err := io.WriteString(w, str)
		if nil != err {
			return err
		}

		expected := len(str)
		actual   := n

		if expected != actual {
			return fmt.Errorf("actual number of byte written (%d) is not what was expected (%d)", actual, expected)
		}

	}
	return nil
}
