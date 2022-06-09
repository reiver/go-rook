package impr_test

import (
	"github.com/reiver/go-rook/internal/token/impr"

	"testing"
)

func TestPull(t *testing.T) {

	const expected string = "/PULL apple-banana-cherry\u0085"

	var sentence impr.Imperation = impr.Pull("apple-banana-cherry")
	var actual string = sentence.String()

	if expected != actual {
		t.Errorf("The actual value is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
