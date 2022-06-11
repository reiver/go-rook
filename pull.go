package rook

import (
	"github.com/reiver/go-rook/internal/token/impr"

	"fmt"
	"io"
)

// PULL writes to ‘w’ the PULL IMPERATION.
//
// I.e.,:
//
//	"/PULL " + object + "\u0085"
func Pull(w io.Writer, object string) error {
	if nil == w {
		return nil
	}

	{
		str := impr.Pull(object).String()

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
