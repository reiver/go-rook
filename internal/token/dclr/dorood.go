package dclr

import (
	"github.com/reiver/go-rook/internal/verbs"
)

// Dorood returns a !DOROOD declaration token.
func Dorood() Declaration {
	return Something(verbs.DOROOD, "")
}
