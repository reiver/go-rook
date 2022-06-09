package dclr

import (
	"github.com/reiver/go-rook/internal/verbs"
)

// Idling returns a !IDLING declaration token.
func Idling() Declaration {
	return Something(verbs.IDLING, "")
}
