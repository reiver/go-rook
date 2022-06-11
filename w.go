package rook

import (
	"github.com/reiver/go-rook/internal/token/impr"

	"fmt"
	"io"
)

// W writes to ‘w’ the W IMPERATION.
//
// I.e.,:
//
//	"/W " + object + "\u0085"
func W(w io.Writer, object string) error {
	if nil == w {
		return nil
	}

	{
		str := impr.W(object).String()

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
