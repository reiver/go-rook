package dclr_test

import (
	"github.com/reiver/go-rook/internal/token/dclr"

	"testing"
)

func TestIdling(t *testing.T) {

	const expected string = "!IDLING\u0085"

	var sentence dclr.Declaration = dclr.Idling()
	var actual string = sentence.String()

	if expected != actual {
		t.Errorf("The actual value is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
